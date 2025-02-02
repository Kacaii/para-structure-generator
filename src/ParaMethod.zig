//! Module ParaMethod provides custom types and functions
//! for structuring PARA Method's Directories.

const std = @import("std");
const testing = std.testing;

/// Provides custom types and functions
/// for structuring PARA Method's Directories.
pub const ParaMethod = struct {
    const Self = @This();

    const ParaDirectory = union(enum) {
        Projects,
        Areas,
        Resources,
        Archive,
    };

    name: ParaDirectory,
    readme_content: []const u8,

    pub fn init(name: Self.ParaDirectory, readme_content: []const u8) Self {
        return .{
            .name = name, //
            .readme_content = readme_content, //
        };
    }

    /// Returns a string containing the number and name of the directory.
    pub fn toString(self: Self) []const u8 {
        return switch (self.name) {
            .Projects => "01 PROJECTS", //       01 Projects/
            .Areas => "02 AREAS", //             02 Areas/
            .Resources => "03 RESOURCES", //     03 Resources/
            .Archive => "04 ARCHIVE", //         04 Archive/
        };
    }
};

/// Stores notes and files for active, time-bound tasks or deliverables.
pub const dir_projects = ParaMethod.init(.Projects,
    \\# 01 PROJECTS
    \\
    \\Stores notes and files for active, time-bound tasks or deliverables.
);

/// Contains ongoing responsibilities or areas of interest.
pub const dir_areas = ParaMethod.init(.Areas,
    \\# 02 AREAS
    \\
    \\Contains ongoing responsibilities or areas of interest.
);

/// Holds general reference materials and reusable templates.
pub const dir_resources = ParaMethod.init(.Resources,
    \\# 03 RESOURCES
    \\
    \\Holds general reference materials and reusable templates.
);

/// Keeps inactive projects and outdated resources for future reference.
pub const dir_archive = ParaMethod.init(.Archive,
    \\# 04 ARCHIVE
    \\
    \\Keeps inactive projects and outdated resources for future reference.
);

test "toString" {
    // 01 Projects
    const projects = ParaMethod.init(.Projects, "");
    try testing.expectEqualStrings("01 PROJECTS", projects.toString());

    // 02 Areas
    const areas = ParaMethod.init(.Areas, "");
    try testing.expectEqualStrings("02 AREAS", areas.toString());

    // 03 Resources
    const resources = ParaMethod.init(.Resources, "");
    try testing.expectEqualStrings("03 RESOURCES", resources.toString());

    // 04 Archive
    const archive = ParaMethod.init(.Archive, "");
    try testing.expectEqualStrings("04 ARCHIVE", archive.toString());
}
