//! Module ParaMethod provides custom types and functions
//! for structuring PARA Method's Directories.

const std = @import("std");
const testing = std.testing;

/// Provides custom types and functions
/// for structuring PARA Method's Directories.
pub const ParaDirectory = struct {
    const Self = @This();

    const NameTag = union(enum) {
        Projects,
        Areas,
        Resources,
        Archive,
    };

    /// Name of the ParaDirectory.
    /// Use dot notation to access its possible values.
    name: NameTag,
    /// Contains a brief description of the directory's purpose.
    readme_content: []const u8,

    pub fn init(name: Self.NameTag, readme_content: []const u8) Self {
        return .{
            .name = name,
            .readme_content = readme_content,
        };
    }
};
