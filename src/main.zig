const std = @import("std");
const ParaMethod = @import("ParaMethod.zig").ParaMethod;

// Storing string in a constant for reusability.
const README_FILE = "README.md";

/// Stores ANSI escape codes for output styling.
const AnsiEscape = struct {
    const reset = "\x1b[0m";
    const green = "\x1b[32m";
    const blue = "\x1b[94m";
};

/// Stores notes and files for active, time-bound tasks or deliverables.
const dir_projects = ParaMethod.init(.Projects,
    \\# 01 PROJECTS
    \\
    \\Stores notes and files for active, time-bound tasks or deliverables.
);

/// Contains ongoing responsibilities or areas of interest.
const dir_areas = ParaMethod.init(.Areas,
    \\# 02 AREAS
    \\
    \\Contains ongoing responsibilities or areas of interest.
);

/// Holds general reference materials and reusable templates.
const dir_resources = ParaMethod.init(.Resources,
    \\# 03 RESOURCES
    \\
    \\Holds general reference materials and reusable templates.
);

/// Keeps inactive projects and outdated resources for future reference.
const dir_archive = ParaMethod.init(.Archive,
    \\# 04 ARCHIVE
    \\
    \\Keeps inactive projects and outdated resources for future reference.
);

/// Storing all necessary directories for iteration.
const para_directories = [4]ParaMethod{
    dir_projects, //    01 Projects/
    dir_areas, //       02 Areas/
    dir_resources, //   03 Resources/
    dir_archive, //     04 Archive/
};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const allocator = gpa.allocator();

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
    const std_out = std.io.getStdOut().writer();

    // Just adding a line feed, nothing fancy.
    try std_out.writeAll("\n");

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

        // Providing feedback for the user.
        switch (i) {
            0 => try std_out.writeAll("┎╴"), //      ┎╴
            else => try std_out.writeAll("┠╴"), //   ┠╴ 
            3 => try std_out.writeAll("┖╴"), //      ┖╴
        }
        try std_out.print("{s} Directory created.\n", .{dir.toString()});

        // Accessing the generated directory.
        var sub_dir = try base_directory.openDir(dir.toString(), .{});
        defer sub_dir.close();

        // Creates and Write contents to README.md file.
        try writeContentToReadME(&sub_dir, dir);

        // Verifies if its in the last iteration.
        if (i != para_directories.len - 1) {
            try std_out.print("┃  ┖╴{s} generated!\n", .{README_FILE});
            try std_out.print("┃  \n", .{});
        } else {
            try std_out.print("   ┖╴{s} generated!\n", .{README_FILE});
        }
    }

    // Script ( hopefully 󱜙 ) completed successfully! 󱁖
    try std_out.writeAll(AnsiEscape.green);
    try std_out.writeAll("\n▒ All done! ▒\n\n");
    try std_out.writeAll(AnsiEscape.reset);
}

/// Creates an README and writes content to it.
fn writeContentToReadME(dir: *std.fs.Dir, para_directory: ParaMethod) !void {
    // Generate a ReadME.md file. 
    const readme_file = try dir.createFile(README_FILE, .{});
    defer readme_file.close();

    // Write content to it. 
    _ = try readme_file.write(para_directory.readme_content);
}
