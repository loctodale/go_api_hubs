using System.Text.Json;
using System.Text.Json.Serialization;

namespace ElsaStudio;

/// <summary>
/// Custom JSON serializer context for WebAssembly compatibility
/// This helps avoid NullabilityInfoContext issues in WASM
/// </summary>
[JsonSerializable(typeof(object))]
[JsonSerializable(typeof(string))]
[JsonSerializable(typeof(int))]
[JsonSerializable(typeof(bool))]
[JsonSerializable(typeof(JsonElement))]
[JsonSerializable(typeof(Dictionary<string, object>))]
[JsonSerializable(typeof(Dictionary<string, string>))]
public partial class WebAssemblyJsonSerializerContext : JsonSerializerContext
{
    /// <summary>
    /// Default options that are WebAssembly-compatible
    /// </summary>
    public static JsonSerializerOptions DefaultOptions => new JsonSerializerOptions
    {
        PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
        DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
        PropertyNameCaseInsensitive = true,
        TypeInfoResolver = Default
    };
}
