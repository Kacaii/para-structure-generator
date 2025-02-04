# PARA Directory Generator

A simple Zig tool to generate a PARA (Projects, Areas, Resources, Archive) directory
structure with README files.

```bash
generate_para /path/to/base/directory
```

If not directory is provided, it will use the current directory.

## What it does

Creates the following file structure:

```txt
├── 01 Projects/
│   └── README.md
├── 02 Areas/
│   └── README.md
├── 03 Resources/
│   └── README.md
└── 04 Archive/
    └── README.md
```

Each README.md contains a brief description of the directory's purpose.

## Example

```bash
generate_para ~/Documents
```

This will create the PARA structure in `~/Documents`.
