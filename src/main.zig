const std = @import("std");
const ParaDirectory = @import("ParaDirectory.zig").ParaDirectory;
const help_file = @embedFile("help.txt");

const README_FILE = "README.md";

const AnsiEscape = struct {
    const DEFAULT_FOREGROUND = "\x1b[0m";
    const GREEN = "\x1b[32m";
};

const stdout = std.io.getStdOut().writer();

/// Storing all necessary directories for iteration.
const para_directories = [4]ParaDirectory{
    dir_projects, //    01 Projects/
    dir_areas, //       02 Areas/
    dir_resources, //   03 Resources/
    dir_archive, //     04 Archive/
};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    const allocator = gpa.allocator();

    // Replace {{green}}
    const help_file_green_replaced = try std.mem.replaceOwned(u8, allocator, help_file, "{{green}}", AnsiEscape.GREEN);
    defer allocator.free(help_file_green_replaced);

    // Replace {{default_foreground}}
    const parsed_help_file = try std.mem.replaceOwned(u8, allocator, help_file_green_replaced, "{{default_foreground}}", AnsiEscape.DEFAULT_FOREGROUND);
    defer allocator.free(parsed_help_file);

    const args = std.process.argsAlloc(allocator) catch |err| {
        std.debug.print("Alloc failed {?}\n", .{err});
        return;
    };
    defer std.process.argsFree(allocator, args);

    // Handle printing help message
    if (args.len > 1 and std.mem.eql(u8, args[1], "help")) {
        try stdout.writeAll(parsed_help_file);
        return;
    }

    var cwd = std.fs.cwd();

    // Uses the provided path as base directory for generate the structure.
    // Defaults to the current one.
    const base_directory = if (args.len == 1) cwd else cwd.openDir(args[1], .{}) catch |err|
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
    };

    try stdout.writeAll("\n");

    // For every item on the para_directories array, generate the respective directory,
    // and write content to its README.md file.
    for (para_directories, 0..) |dir, i| {
        const dir_name_tag = switch (dir.name) {
            .Projects => "01 PROJECTS", //       01 Projects/
            .Areas => "02 AREAS", //             02 Areas/
            .Resources => "03 RESOURCES", //     03 Resources/
            .Archive => "04 ARCHIVE", //         04 Archive/
        };

        // Creating a sub_path inside the directory provided.
        base_directory.makeDir(dir_name_tag) catch |err| switch (err) {
            error.PathAlreadyExists => {
                std.log.err("The directory already exists: {s}\n", .{dir_name_tag});
                return;
            },
            else => return err,
        };

        // Printing the file tree.
        switch (i) {
            // First directory
            0 => try stdout.print("┎╴{s} Directory created.\n", .{dir_name_tag}),

            // Second and Third Directories
            else => try stdout.print("┠╴{s} Directory created.\n", .{dir_name_tag}),

            // Last directory
            (para_directories.len - 1) => try stdout.print("┖╴{s} Directory created.\n", .{dir_name_tag}),
        }

        // Accessing the generated directory.
        var sub_dir = try base_directory.openDir(dir_name_tag, .{});
        defer sub_dir.close();

        // Creates and Write contents to README.md file.
        try writeReadME(&sub_dir, dir);

        // Verifying if its the last interation.
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
    try stdout.writeAll(AnsiEscape.DEFAULT_FOREGROUND);
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
const dir_projects = ParaDirectory.init(.Projects,
    \\# 01 PROJECTS
    \\
    \\Stores notes and files for active, time-bound tasks or deliverables.
);

/// Contains ongoing responsibilities or areas of interest.
const dir_areas = ParaDirectory.init(.Areas,
    \\# 02 AREAS
    \\
    \\Contains ongoing responsibilities or areas of interest.
);

/// Holds general reference materials and reusable templates.
const dir_resources = ParaDirectory.init(.Resources,
    \\# 03 RESOURCES
    \\
    \\Holds general reference materials and reusable templates.
);

/// Keeps inactive projects and outdated resources for future reference.
const dir_archive = ParaDirectory.init(.Archive,
    \\# 04 ARCHIVE
    \\
    \\Keeps inactive projects and outdated resources for future reference.
);
