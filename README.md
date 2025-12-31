# Appstreamfile
A Dockerfile for AWS Image Builders Built in Go. A declarative spec for AWS AppStream images. Appstreamfile defines everything that goes into an image, while your automation system handles execution. 

# Why
AWS AppStream provides a CLI, but it only handles high-level image operations like updating stack applications or building images. The real pain is preparing the image itself. Its the installers, custom scripts, environment tweaks, catalog entries, and session scripts. Doing this by hand or via plain scripting is messy and impossible to audit cleanly.

**Appstreamfile** is the missing piece. It acts like a Dockerfile for AppStream image builders. Install it inside the image builder instance with your favorite automation of choice, point it to a config file, and execute. Your image becomes reproducible, reviewable, and predictable.

# Features
1. Single config file describing the image
2. Scripted installers
3. App catalog configuration
4. System/user session scripts
5. File provisioning
6. Declarative image metadata
7. S3 Backend Support

# Planned Features
1. Registry Configs
2. Scaffolded startup Apps - Customise your app startup
3. Git backend support
4. State persistence within image builders 

# Installation
## Windows Installation
You can install and run **AppStreamFile** without cloning the repository.  
The script detects your system architecture and downloads the correct binary automatically.

```powershell
iex (iwr https://raw.githubusercontent.com/aslamcodes/appstreamfile/main/win_install.ps1 -UseBasicParsing); Install-AppStreamFile v0.2.0
```

## Linux
***Supported in future releases***

# Usage
Appstreamfile supports fetching configuration from local filesystem or S3.

### Local Source
Use **local** backend source when the config file is within the image builder.
```sh
    appstreamfile -source local -location config.yaml
```

### S3 Source
Use **s3** backend source to fetch config from an S3 bucket.
*Note: Requires an AWS profile named `appstream_machine_role` to be configured.*
```sh
    appstreamfile -source s3 -bucket my-bucket -key path/to/config.yaml
```

### Flags
| Flag | Description | Required For |
|------|-------------|--------------|
| `-source` | Configuration source (`local` or `s3`) | All |
| `-location` | Path to local config file | `local` source |
| `-bucket` | S3 bucket name | `s3` source |
| `-key` | S3 object key | `s3` source |
| `-version-id` | S3 object version ID (optional) | `s3` source |

# Configuration Reference

The configuration file is a YAML file with the following sections:

### `platform`
Operating system platform (e.g., `windows`).

### `installers`
List of installation scripts to run.
- `executable`: The shell/executable to run the script (e.g., `powershell`).
- `installScript`: The actual script content.

### `catalog`
List of applications to add to the AppStream catalog.
- `name`: Internal name of the application.
- `path`: Absolute path to the executable.
- `display_name`: Name shown to users.
- `parameters`: Launch parameters.
- `icon_path`: Path to the icon file.
- `working_dir`: Working directory for the application.

### `files`
List of files to create on the system.
- `path`: Destination path.
- `content`: Content of the file.

### `session_scripts`
Configuration for session scripts (start and termination).
Contains `session_start` and `session_termination` blocks, each with:
- `waitingTime`: Max time to wait for scripts to complete (in seconds).
- `executables`: List of scripts to run.
  - `context`: `system` or `user`.
  - `filename`: Path to the script.
  - `arguments`: Arguments to pass.
  - `s3LogEnabled`: Boolean to enable logging to S3.

### `image`
Metadata for the resulting AppStream image.
- `name`: Image name.
- `display_name`: Display name.
- `description`: Description.
- `enable_dynamic_app_catalog`: Boolean.
- `use_latest_agent_version`: Boolean.
- `tags`: List of tags (`key:value`).
- `dry_run`: If true, simulates operations without making changes.

# Sample config
```yaml
platform: "windows"
installers:
  - executable: "powershell"
    installScript: |
      Write-Host "Hello World"

  - executable: "powershell"
    installScript: |
      echo "Setting up environment"

catalog:
  - name: "Notepad"
    path: "C:\\Windows\\System32\\notepad.exe"
    display_name: "Notepad"
    parameters: ""
    icon_path: "C:\\Windows\\System32\\notepad.exe"
    working_dir: "C:\\Windows\\System32"

files:
  - path: "C:\\AppStream\\Scripts\\Start-System.ps1"
    content: |
      Write-EventLog -LogName Application -Source AppStream -EventID 100 -Message "System session start"

  - path: "C:\\AppStream\\Scripts\\Start-User.ps1"
    content: |
      Write-Host "User profile initialization"

  - path: "C:\\AppStream\\Scripts\\End-System.ps1"
    content: |
      Write-EventLog -LogName Application -Source AppStream -EventID 200 -Message "System cleanup"

  - path: "C:\\AppStream\\Scripts\\End-User.ps1"
    content: |
      Write-Host "User session cleanup"

session_scripts:
  session_start:
    executables:
      - context: "system"
        filename: "C:\\AppStream\\Scripts\\Start-System.ps1"
        arguments: "-Init"
        s3LogEnabled: true
      - context: "user"
        filename: "C:\\AppStream\\Scripts\\Start-User.ps1"
        arguments: "-UserSetup"
        s3LogEnabled: true
    waitingTime: 60

  session_termination:
    executables:
      - context: "system"
        filename: "C:\\AppStream\\Scripts\\End-System.ps1"
        arguments: "-Cleanup"
        s3LogEnabled: true
      - context: "user"
        filename: "C:\\AppStream\\Scripts\\End-User.ps1"
        arguments: "-CleanupUser"
        s3LogEnabled: true
    waitingTime: 60

image:
  name: "example_image"
  display_name: "example image"
  description: "example image"
  enable_dynamic_app_catalog: true
  use_latest_agent_version: false
  tags:
    - team:infra
    - env:dev
  dry_run: false
```
