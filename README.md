# Cadagen FS Microservice

A dedicated microservice within the Cadagen ecosystem for performing low-level filesystem operations. This service is designed to be called by backend controllers to provide a standardized, cross-platform interface for file access via CLI.

## Execution Model

The service is a stateless CLI binary. It receives commands via standard arguments and writes a single JSON object to stdout.

### Response Schema

Every execution results in a JSON object with a type discriminator.

**Success Response**

```json
{
  "type": "ok",
  "data": T
}
```

_Note: data type T depends on the command executed._

**Error Response**

```json
{
  "type": "error",
  "message": "Detailed error description"
}
```

---

## Commands

### read

Reads the content of a file. Supports optional line-range limiting to handle large text files or logs efficiently.

**Arguments:**
`read -start=[start_line] -end=[end_line] <path>`

- **path**: Absolute or relative path to the file.
- **start_line** (optional): The inclusive line number to begin reading (1-indexed).
- **end_line** (optional): The inclusive line number to end reading.

**Returns:** `string` (The requested lines of the file).

### ls

Lists the contents of a directory.

**Arguments:**
`ls <path>`

- **path**: Path to the target directory.

**Returns:**

```json
[
  { "type": "file", "name": "example.txt" },
  { "type": "dir", "name": "subdirectory" }
]
```

---

## Technical Integration

### Error Handling

The backend should check the `type` field before attempting to access `data`.

### Security Note

The backend calling this service is responsible for path sanitization. Ensure that user-inputted paths are validated to prevent directory traversal attacks (e.g., `../../etc/passwd`).

---

## Examples

**Read specific line range (e.g., lines 10 through 50):**

```bash
./cadagen-fs read -start=10 -end=50 ./logs/app.log
```

**List directory:**

```bash
./cadagen-fs ls ./uploads/user_1
```
