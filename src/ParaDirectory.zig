//! Module ParaDirectory provides custom types and functions
//! for structuring PARA Method's Directories.

const std = @import("std");
const testing = std.testing;

/// ParaDirectory defines a directory in the PARA structure with a name and description.
pub const ParaDirectory = struct {
    /// Name of the Directory.
    name: union(enum) { Projects, Areas, Resources, Arquive },

    /// Content to be written to the README.md file inside every PARA directory.
    readme_content: []const u8,

    /// A constant reference to itself.
    const self = @This();

    // Returns a string containing the number and name of the directory.
    pub fn getName(s: self) []const u8 {
        return switch (s.name) {
            .Projects => "01 PROJECTS",
            .Areas => "02 AREAS",
            .Resources => "03 RESOURCES",
            .Arquive => "04 ARQUIVE",
        };
    }
};

test "getName" {
    // 01 Projects
    const projects = ParaDirectory{
        .name = .Projects,
        .readme_content = "",
    };
    try testing.expectEqualStrings("01 PROJECTS", projects.getName());

    // 02 Areas
    const areas = ParaDirectory{
        .name = .Areas,
        .readme_content = "",
    };
    try testing.expectEqualStrings("02 AREAS", areas.getName());

    // 03 Resources
    const resources = ParaDirectory{
        .name = .Resources,
        .readme_content = "",
    };
    try testing.expectEqualStrings("03 RESOURCES", resources.getName());

    // 04 Arquive
    const arquive = ParaDirectory{
        .name = .Arquive,
        .readme_content = "",
    };
    try testing.expectEqualStrings("04 ARQUIVE", arquive.getName());
}
