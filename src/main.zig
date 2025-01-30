const std = @import("std");
const ParaDirectory = @import("ParaDirectory.zig").ParaDirectory;

// Storing string in a constant for reusability.
const README_FILE = "README.md";

/// Stores notes and files for active, time-bound tasks or deliverables.
const dir_projects = ParaDirectory{ //
    .name = .Projects,
    .readme_content =
    \\# 01 PROJECTS
    \\
    \\Stores notes and files for active, time-bound tasks or deliverables.
};

/// Contains ongoing responsibilities or areas of interest.
const dir_areas = ParaDirectory{ //
    .name = .Areas,
    .readme_content =
    \\# 02 AREAS
    \\
    \\Contains ongoing responsibilities or areas of interest.
};

/// Holds general reference materials and reusable templates.
const dir_resources = ParaDirectory{ //
    .name = .Resources,
    .readme_content =
    \\# 03 RESOURCES
    \\
    \\Holds general reference materials and reusable templates.
};

/// Keeps inactive projects and outdated resources for future reference.
const dir_archive = ParaDirectory{ //
    .name = .Archive,
    .readme_content =
    \\# 04 ARCHIVE
    \\
    \\Keeps inactive projects and outdated resources for future reference.
};

/// Storing all necessary directories for iteration.
const para_directories = [4]ParaDirectory{
    dir_projects, //    01 Projects/
    dir_areas, //       02 Areas/
    dir_resources, //   03 Resources/
    dir_archive, //     04 Archive/
};

/// This is the entry point of the program.
pub fn main() !void {
    // Getting current working directory.
    var cwd = std.fs.cwd();

    // Just adding a line feed, nothing fancy.
    std.debug.print("\n", .{});

    // For every item on the para_directories array, generate the respective directory,
    // and write content to its  ReadME.md file.
    for (para_directories, 0..) |dir, i| {
        // Generate directory 
        try generateParaDirectory(&cwd, dir);

        // Printing the file tree.
        switch (i) {
            0 => std.debug.print("┎╴", .{}), //      ┎╴ First Directory/
            else => std.debug.print("┠╴", .{}), //   ┠╴ Middle Directories/
            3 => std.debug.print("┖╴", .{}), //      ┖╴ Last Directory/
        }

        // Feedback for the user.
        std.debug.print("{s} Directory created.\n", .{dir.getName()});

        // Creates and Write contents to README.md file.
        try writeContentToReadME(&cwd, dir);

        // Check for last directory. 
        if (i == para_directories.len - 1) {
            // If its the last one, the file tree ends.
            std.debug.print("    ┖╴{s} generated!\n", .{README_FILE});
        } else {
            // If its not the last one, the tree continues.
            std.debug.print("┃   ┖╴{s} generated!\n", .{README_FILE});
            std.debug.print("┃   \n", .{});
        }
    }

    // Program ( hopefully 󱜙 ) completed successfully! 󱁖
    std.debug.print("\n▒ All done! ▒\n\n", .{});
}

/// Takes an directory and creates a new one using the provided ParaDirectory.
fn generateParaDirectory(dir: *std.fs.Dir, para_directory: ParaDirectory) std.fs.Dir.MakeError!void {
    // Generating a sub_path inside the directory provided.
    dir.makeDir(para_directory.getName()) catch |err| switch (err) {
        error.PathAlreadyExists => {
            // Add error handling for this case since its a more common one.
            std.log.err("The directory already exists: {s}\n", .{para_directory.getName()});
            std.process.exit(1); // Finishing the program.
        },

        else => return err,
    };
}

/// Creates an README and writes content to it.
fn writeContentToReadME(dir: *std.fs.Dir, para_directory: ParaDirectory) !void {
    var sub_dir = try dir.openDir(para_directory.getName(), .{});
    defer sub_dir.close();

    // Generate a ReadME.md file. 
    const readme_file = try sub_dir.createFile(README_FILE, .{});
    defer readme_file.close();

    // Write content to it. 
    _ = try readme_file.write(para_directory.readme_content);
}
