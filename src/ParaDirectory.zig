const std = @import("std");
const testing = std.testing;

pub const ParaDirectory = struct {
    category: union(enum) { Projects, Areas, Resources, Arquive },
    readme_content: []const u8,

    const self = @This();

    pub fn getCategory(s: self) []const u8 {
        return switch (s.category) {
            .Projects => "01 Projects",
            .Areas => "02 Areas",
            .Resources => "03 Resources",
            .Arquive => "04 Arquive",
        };
    }
};

test "getCategory" {
    // 01 Projects
    const projects = ParaDirectory{
        .category = .Projects,
        .readme_content = "",
    };
    try testing.expectEqualStrings("01 Projects", projects.getCategory());

    // 02 Areas
    const areas = ParaDirectory{
        .category = .Areas,
        .readme_content = "",
    };
    try testing.expectEqualStrings("02 Areas", areas.getCategory());

    // 03 Resources
    const resources = ParaDirectory{
        .category = .Resources,
        .readme_content = "",
    };
    try testing.expectEqualStrings("03 Resources", resources.getCategory());

    // 04 Arquive
    const arquive = ParaDirectory{
        .category = .Arquive,
        .readme_content = "",
    };
    try testing.expectEqualStrings("04 Arquive", arquive.getCategory());
}
