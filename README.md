# PowerShell Script Obfuscator

This tool is a PowerShell script obfuscator implemented in Go. It provides various obfuscation techniques to make PowerShell scripts more challenging to analyze and reverse-engineer.

## Features

- **Base64 Encoding**: Encode the PowerShell script in Base64.
- **Layered Encoding**: Apply multiple rounds of Base64 encoding with additional transformations.
- **String Splitting and Reassembly**: Split the encoded script into parts and reassemble it.
- **NOP Insertion**: Insert random No-Operation (NOP) commands to confuse static analysis.
- **Control Flow Obfuscation**: Add fake branching to obscure the script's control flow.
- **Randomized Names**: Replace variable and function names with randomized strings.

## Installation

1. **Clone the Repository**:
   ```sh
   https://github.com/Algorora/PowerShell-Obufscator.git 
   cd Powershell-Obfuscator
   ```
## Build the Executable

Ensure you have Go installed. Build the Go script into an executable:
   ```sh
   go build obfuscator.go
   ```

## Usage

To obfuscate a PowerShell script, use the following command:
- **Obufscate a script**:
   ```sh
   ./obfuscator -script /path/to/your/script.ps1
   ```
- **Save output to a file**:
   ```sh
   ./obfuscator -script /path/to/your/script.ps1 > obfuscated_script.ps1
   ```

## Disclaimer

To the maximum extent permitted by applicable law, I and/or affiliates whom this repo is sourced and or submitted content to this repo, shall not be liable for any indirect, incidental, special, consequential or punitive damages, or any loss of profits or revenue, whether incurred directly or indirectly, or any loss of data, use, goodwill, or other intangible losses, resulting from:

👉 (i) your access to this resource and/or inability to access this resource; 👉 (ii) any conduct or content of any third party referenced by this resource, including, any offensive or illegal conduct or other users or third parties; 👉 (iii) any content obtained from this resource or any of its resources
