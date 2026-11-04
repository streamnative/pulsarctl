#!/usr/bin/env bash
set -euo pipefail

version=${1:-v0.69.3}
case "$(uname -s)" in
  Linux) os=Linux ;;
  Darwin) os=macOS ;;
  *) echo "Unsupported operating system" >&2; exit 1 ;;
esac
case "$(uname -m)" in
  x86_64) arch=64bit ;;
  aarch64|arm64) arch=ARM64 ;;
  *) echo "Unsupported architecture" >&2; exit 1 ;;
esac

archive="trivy_${version#v}_${os}-${arch}.tar.gz"
checksums="trivy_${version#v}_checksums.txt"
release_url="https://github.com/aquasecurity/trivy/releases/download/${version}"
download_dir=$(mktemp -d)
trap 'rm -rf "$download_dir"' EXIT
install_dir="${RUNNER_TEMP:-${TMPDIR:-/tmp}}/trivy-bin"

for asset in "$archive" "$checksums"; do
  curl --fail --show-error --location --retry 3 --retry-all-errors \
    --connect-timeout 15 --max-time 180 \
    --output "$download_dir/$asset" "$release_url/$asset"
done

checksum=$(awk -v archive="$archive" '$2 == archive {print $1}' "$download_dir/$checksums")
if [[ ! "$checksum" =~ ^[[:xdigit:]]{64}$ ]]; then
  echo "Missing or invalid checksum for $archive" >&2
  exit 1
fi
(
  cd "$download_dir"
  if command -v sha256sum >/dev/null 2>&1; then
    printf '%s  %s\n' "$checksum" "$archive" | sha256sum --check -
  else
    printf '%s  %s\n' "$checksum" "$archive" | shasum -a 256 --check -
  fi
)

mkdir -p "$install_dir"
tar -xzf "$download_dir/$archive" -C "$install_dir" trivy
"$install_dir/trivy" --version
if [[ -n "${GITHUB_PATH:-}" ]]; then
  printf '%s\n' "$install_dir" >> "$GITHUB_PATH"
fi
