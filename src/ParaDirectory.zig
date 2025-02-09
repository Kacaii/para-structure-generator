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
    /// Use dot notation to access its possible values.
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
};
