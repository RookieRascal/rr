const std = @import("std");

pub fn main() !void {
    const file_path = "index.txt";
    const content = "Index data for storage";

    // Create or open the file
    var file = try std.fs.cwd().createFile(file_path, .{});
    defer file.close();

    // Write content to the file
    try file.writeAll(content);

    // Use the correct format specifier for strings
    std.debug.print("File {s} written.\n", .{file_path});
}
