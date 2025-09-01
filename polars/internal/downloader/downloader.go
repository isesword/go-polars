package downloader

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	defaultVersion = "v0.0.26"
	defaultRepo    = "jordandelbar/go-polars"
)

// Options configures Ensure.
type Options struct {
	Version    string
	Force      bool
	SkipVerify bool
	BinDir     string
	BaseURL    string
	HTTPClient *http.Client
}

// Ensure downloads the correct static library for the host platform if needed.
func Ensure(opts Options) error {
	if opts.Version == "" {
		opts.Version = versionFromEnv()
	}
	if opts.BinDir == "" {
		opts.BinDir = "polars/bin"
	}
	if opts.BaseURL == "" {
		opts.BaseURL = fmt.Sprintf("https://github.com/%s/releases/download", defaultRepo)
	}
	client := opts.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: 2 * time.Minute}
	}

	if err := os.MkdirAll(opts.BinDir, 0o755); err != nil {
		return fmt.Errorf("create bin dir: %w", err)
	}

	platform, err := hostPlatform()
	if err != nil {
		return err
	}

	binName := binaryName()
	targetPath := filepath.Join(opts.BinDir, binName)
	statePath := filepath.Join(opts.BinDir, ".lib_version")

	if !opts.Force && fileMatchesVersion(targetPath, statePath, opts.Version) {
		return nil
	}

	assetName := fmt.Sprintf("libpolars_go-%s-%s.a", platform, opts.Version)
	assetURL := fmt.Sprintf("%s/%s/%s", opts.BaseURL, opts.Version, assetName)
	checksumURL := assetURL + ".sha256"

	fmt.Printf("⬇️  downloading %s\n", assetName)
	tmpFile, err := os.CreateTemp("", "polars-*.a")
	if err != nil {
		return fmt.Errorf("tmp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if err := downloadTo(client, assetURL, tmpFile); err != nil {
		return err
	}

	if !opts.SkipVerify {
		checksum, err := fetchChecksum(client, checksumURL)
		if err != nil {
			return err
		}
		if err := verifyFile(tmpFile.Name(), checksum); err != nil {
			return err
		}
	}

	if err := copyFile(tmpFile.Name(), targetPath); err != nil {
		return err
	}

	if err := os.WriteFile(statePath, []byte(opts.Version), 0o644); err != nil {
		return fmt.Errorf("write version state: %w", err)
	}

	fmt.Printf("✅ %s ready in %s\n", binName, opts.BinDir)
	return nil
}

func versionFromEnv() string {
	if v := os.Getenv("POLARS_BINARY_VERSION"); v != "" {
		return v
	}
	return defaultVersion
}

func hostPlatform() (string, error) {
	switch runtime.GOOS {
	case "linux":
		if runtime.GOARCH == "amd64" {
			return "linux-amd64", nil
		}
	case "darwin":
		switch runtime.GOARCH {
		case "arm64":
			return "darwin-arm64", nil
		case "amd64":
			return "darwin-amd64", nil
		}
	case "windows":
		if runtime.GOARCH == "amd64" {
			return "windows-amd64", nil
		}
	}
	return "", fmt.Errorf("unsupported platform %s/%s", runtime.GOOS, runtime.GOARCH)
}

func binaryName() string {
	if runtime.GOOS == "windows" {
		return "polars_go.lib"
	}
	return "libpolars_go.a"
}

func fileMatchesVersion(binPath, statePath, version string) bool {
	if _, err := os.Stat(binPath); err != nil {
		return false
	}
	data, err := os.ReadFile(statePath)
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(data)) == version
}

func downloadTo(client *http.Client, url string, dst *os.File) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("download %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("release asset missing: %s", url)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download %s: %s", url, resp.Status)
	}

	if _, err := io.Copy(dst, resp.Body); err != nil {
		return fmt.Errorf("write temp file: %w", err)
	}

	if _, err := dst.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("rewind temp file: %w", err)
	}

	return nil
}

func fetchChecksum(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("download checksum: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("checksum request failed: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	if !scanner.Scan() {
		return "", errors.New("empty checksum file")
	}
	line := scanner.Text()
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return "", errors.New("invalid checksum format")
	}
	return fields[0], nil
}

func verifyFile(path, expected string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open for verify: %w", err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return fmt.Errorf("hash file: %w", err)
	}

	sum := hex.EncodeToString(h.Sum(nil))
	if !strings.EqualFold(sum, expected) {
		return fmt.Errorf("checksum mismatch: have %s want %s", sum, expected)
	}
	return nil
}

func copyFile(src, dst string) error {
	input, err := os.Open(src)
	if err != nil {
		return err
	}
	defer input.Close()

	tmp, err := os.CreateTemp(filepath.Dir(dst), "polars-lib-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())

	if _, err := io.Copy(tmp, input); err != nil {
		tmp.Close()
		return err
	}

	if err := tmp.Close(); err != nil {
		return err
	}

	if err := os.Rename(tmp.Name(), dst); err != nil {
		return err
	}

	return nil
}
