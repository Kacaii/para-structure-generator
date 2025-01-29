const std = @import("std");
const ParaDirectory = @import("ParaDirectory.zig").ParaDirectory;

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
const para_directories = [4]ParaDirectory{ //
    dir_projects, //    01 Projects/
    dir_areas, //       02 Areas/
    dir_resources, //   03 Resources/
    dir_archive, //     04 Archive/
};

/// This is the entry point of the program.
pub fn main() !void {
    // Getting current working directory.
    const cwd = std.fs.cwd();

    // Just adding a line feed, nothing fancy.
    std.debug.print("\n", .{});

    // For every item on the para_directories array,
    // generate the respective directory, and write content to is
    // ReadME file.

    for (para_directories, 0..) |dir, i| {
        // Generate directory 
        cwd.makeDir(dir.getName()) catch |err| switch (err) {
            error.PathAlreadyExists => {
                std.debug.print("Directory already exists!\n\n", .{});
                std.process.exit(1);
            },
            else => return err,
        };

        // Drawing the file tree.
        switch (i) {
            0 => std.debug.print("┎╴", .{}),
            else => std.debug.print("┠╴", .{}),
            3 => std.debug.print("┖╴", .{}),
        }

        std.debug.print("{s} Directory created.\n", .{dir.getName()});

        // Open it. 
        var sub_dir = try cwd.openDir(dir.getName(), .{});
        defer sub_dir.close();

        // Generate a ReadME.md file. 
        const file = try sub_dir.createFile("ReadME.md", .{});
        defer file.close();

        // Write content to it. 
        _ = try file.write(dir.readme_content);

        // Check for last directory. 
        if (i == para_directories.len - 1) {
            // If its the last one, the file tree ends.
            std.debug.print("    ┖╴ReadMe.md generated!\n", .{});
        } else {
            // If its not the last one, the tree continues.
            std.debug.print("┃   ┖╴ReadMe.md generated!\n", .{});
            std.debug.print("┃   \n", .{});
        }
    }

    // Program (probably) completed successfully! 󱁖
    std.debug.print("\n▒ All done! ▒\n\n", .{});
}
