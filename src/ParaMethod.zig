//! Module ParaMethod provides custom types and functions
//! for structuring PARA Method's Directories.

const std = @import("std");
const testing = std.testing;

/// Provides custom types and functions
/// for structuring PARA Method's Directories.
pub const ParaDirectory = struct {
    const Self = @This();

    /// Stores all four possible values of a ParaDirectory.
    const NameTag = union(enum) {
        /// Stores notes and files for active, time-bound tasks or deliverables.
        Projects,
        /// Contains ongoing responsibilities or areas of interest.
        Areas,
        /// Holds general reference materials and reusable templates.
        Resources,
        /// Keeps inactive projects and outdated resources for future reference.
        Archive,
    };

    /// Name of the ParaDirectory.
    /// Use an anonymous struct 󰅪 to access its possible values.
    name: NameTag,
    /// Contains a brief description of the directory's purpose.
    readme_content: []const u8,

    /// Returns an new instance of a ParaMethod Directory
    pub fn init(name: Self.NameTag, readme_content: []const u8) Self {
        return .{
            .name = name,
            .readme_content = readme_content,
        };
    }

    /// Returns a string containing the number and name of the directory.
    pub fn toString(s: Self) []const u8 {
        return switch (s.name) {
            .Projects => "01 PROJECTS", //       01 Projects/
            .Areas => "02 AREAS", //             02 Areas/
            .Resources => "03 RESOURCES", //     03 Resources/
            .Archive => "04 ARCHIVE", //         04 Archive/
        };
    }
};

test "toString" {
    // 01 Projects
    const projects = ParaDirectory.init(.Projects, "");
    try testing.expectEqualStrings("01 PROJECTS", projects.toString());

    // 02 Areas
    const areas = ParaDirectory.init(.Areas, "");
    try testing.expectEqualStrings("02 AREAS", areas.toString());

    // 03 Resources
    const resources = ParaDirectory.init(.Resources, "");
    try testing.expectEqualStrings("03 RESOURCES", resources.toString());

    // 04 Archive
    const archive = ParaDirectory.init(.Archive, "");
    try testing.expectEqualStrings("04 ARCHIVE", archive.toString());
}
