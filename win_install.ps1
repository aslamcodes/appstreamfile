param(
  [string]$Version
)

function Install-AppStreamFile {
  param([string]$Version)

  if (-not $Version) {
    throw "Usage: iwr ... | iex -Version v0.2.0"
  }

  $repo    = "aslamcodes/appstreamfile"
  $baseUrl = "https://github.com/$repo/releases/download/$Version"
  $dir     = "$env:TEMP\appstreamfile"

  if (!(Test-Path $dir)) {
    New-Item -ItemType Directory -Path $dir | Out-Null
  }

  $arch = switch ($env:PROCESSOR_ARCHITECTURE) {
    "AMD64" { "x86_64" }
    "x86"   { "i386" }
    "ARM64" { "arm64" }
    default { throw "Unsupported architecture" }
  }

  $file = "appstreamfile_Windows_$arch.zip"
  $url  = "$baseUrl/$file"
  $zip  = Join-Path $dir $file

  Invoke-WebRequest $url -OutFile $zip
  Expand-Archive $zip $dir -Force
  Remove-Item $zip

  $exe = Join-Path $dir "appstreamfile.exe"
  if (!(Test-Path $exe)) { throw "Binary not found" }

  & $exe --help
}

Install-AppStreamFile -Version $Version
