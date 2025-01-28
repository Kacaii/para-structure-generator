const std = @import("std");
const ParaDirectory = @import("ParaDirectory.zig").ParaDirectory;

const projects = ParaDirectory{
    //
    .name = .Projects,
    .readme_content =
    \\# 01 PROJECTS
    \\
    \\Stores notes and files for active, time-bound tasks or deliverables.
};

const areas = ParaDirectory{
    //
    .name = .Areas,
    .readme_content =
    \\ # 02 AREAS
    \\
    \\Contains ongoing responsibilities or areas of interest.
};

const resources = ParaDirectory{
    //
    .name = .Resources,
    .readme_content =
    \\# 03 RESOURCES
    \\
    \\Holds general reference materials and reusable templates.
};

const arquive = ParaDirectory{
    //
    .name = .Resources,
    .readme_content =
    \\# 04 ARQUIVE
    \\
    \\Keeps inactive projects and outdated resources for future reference.
};

pub fn main() !void {
    const directories = [4]ParaDirectory{
        //
        projects,
        areas,
        resources,
        arquive,
    };
    const cwd = std.fs.cwd();

    for (directories) |dir| {
        try cwd.makeDir(dir.getName());
        const sub_dir = try cwd.openDir(dir.getName(), .{});
        const file = try sub_dir.createFile("ReadME.md", .{});
        try file.write(dir.readme_content);
    }
}
