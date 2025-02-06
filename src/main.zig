const std = @import("std");
const ParaDirectory = @import("ParaMethod.zig").ParaDirectory;

// Storing string in a constant for reusability.
const README_FILE = "README.md";

/// Stores ANSI escape codes for output styling.
const AnsiEscape = struct {
    const RESET = "\x1b[0m";
    const GREEN = "\x1b[32m";
};

/// Storing all necessary directories for iteration.
const para_directories = [4]ParaDirectory{
    dir_projects, //    01 Projects/
    dir_areas, //       02 Areas/
    dir_resources, //   03 Resources/
    dir_archive, //     04 Archive/
};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit(); // Freeing resources and detecting leaks.

    const allocator = gpa.allocator();

    // Storing command-line args in an array.
    const args = std.process.argsAlloc(allocator) catch |err| {
        std.debug.print("Alloc failed {?}\n", .{err});
        return;
    };
    defer std.process.argsFree(allocator, args);

    // Open current directory.
    var cwd = std.fs.cwd();

    // Uses the provided path as base directory for generate the structure.
    // Defaults to the current one.
    const base_directory = if (args.len == 1) cwd else cwd.openDir(args[1], .{}) catch |err| {

        // Handling specific errors.
        switch (err) {
            error.NotDir => {
                std.log.err("Provided Path needs to be a Directory.", .{});
                return;
            },
            error.FileNotFound => {
                std.log.err("Provided Path does not exist.", .{});
                return;
            },
            else => return err, // Returning unexpected error.
        }
    };

    // Get standard output for providing feedback to user.
    const stdout = std.io.getStdOut().writer();

    // Just adding a line feed, nothing fancy.
    try stdout.writeAll("\n");

    // For every item on the para_directories array, generate the respective directory,
    // and write content to its  ReadME.md file.
    for (para_directories, 0..) |dir, i| {

        // Creating a sub_path inside the directory provided.
        base_directory.makeDir(dir.toString()) catch |err| switch (err) {
            error.PathAlreadyExists => {
                std.log.err("The directory already exists: {s}\n", .{dir.toString()});
                return;
            },
            else => return err,
        };

        // Printing the file tree
        switch (i) {
            // First directory
            0 => try stdout.print("┎╴{s} Directory created.\n", .{dir.toString()}),

            // Middle directories
            else => try stdout.print("┠╴{s} Directory created.\n", .{dir.toString()}),

            // Last directory
            (para_directories.len - 1) => try stdout.print("┖╴{s} Directory created.\n", .{dir.toString()}),
        }

        // Accessing the generated directory.
        var sub_dir = try base_directory.openDir(dir.toString(), .{});
        defer sub_dir.close();

        // Creates and Write contents to README.md file.
        try writeReadME(&sub_dir, dir);

        // Verifies if its in the last iteration.
        if (i != para_directories.len - 1) {
            try stdout.print("┃  ┖╴{s} generated!\n", .{README_FILE});
            try stdout.print("┃  \n", .{});
        } else {
            try stdout.print("   ┖╴{s} generated!\n", .{README_FILE});
        }
    }

    // Script ( hopefully 󱜙 ) completed successfully! 󱁖
    try stdout.writeAll(AnsiEscape.GREEN);
    try stdout.writeAll("\n▒ All done! ▒\n\n");
    try stdout.writeAll(AnsiEscape.RESET);
}

/// Creates an README and writes content to it.
fn writeReadME(dir: *std.fs.Dir, para_directory: ParaDirectory) !void {
    // Generate a ReadME.md file. 
    const readme_file = try dir.createFile(README_FILE, .{});
    defer readme_file.close();

    // Write content to it. 
    _ = try readme_file.write(para_directory.readme_content);
}

/// Stores notes and files for active, time-bound tasks or deliverables.
pub const dir_projects = ParaDirectory.init(.Projects,
    \\# 01 PROJECTS
    \\
    \\Stores notes and files for active, time-bound tasks or deliverables.
);

/// Contains ongoing responsibilities or areas of interest.
pub const dir_areas = ParaDirectory.init(.Areas,
    \\# 02 AREAS
    \\
    \\Contains ongoing responsibilities or areas of interest.
);

/// Holds general reference materials and reusable templates.
pub const dir_resources = ParaDirectory.init(.Resources,
    \\# 03 RESOURCES
    \\
    \\Holds general reference materials and reusable templates.
);

/// Keeps inactive projects and outdated resources for future reference.
pub const dir_archive = ParaDirectory.init(.Archive,
    \\# 04 ARCHIVE
    \\
    \\Keeps inactive projects and outdated resources for future reference.
);
