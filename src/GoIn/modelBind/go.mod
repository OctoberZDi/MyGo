module GoIn/modelBind

go 1.16

//要将请求正文绑定到类型，请使用模型绑定。我们目前支持 JSON、XML、YAML 和标准表单值 (foo=bar&boo=baz) 的绑定。
//
//Gin 使用 go-playground/validator/v10 进行验证。在此处查看有关标签使用的完整文档。
//
//请注意，您需要在所有要绑定的字段上设置相应的绑定标签。例如，从 JSON 绑定时，设置 json:"fieldname"。
//
//另外，Gin 提供了两组绑定方法：

//类型 - 必须绑定 MustBind
//方法 - 绑定、BindJSON、BindXML、BindQuery、BindYAML、BindHeader
//行为 - 这些方法在幕后使用 MustBindWith。如果存在绑定错误，则请求将通过 c.AbortWithError(400, err).SetType(ErrorTypeBind) 中止。这会将响应状态代码设置为 400，并将 Content-Type 标头设置为 text/plain；字符集=utf-8。请注意，如果您尝试在此之后设置响应代码，则会导致警告 [GIN-debug] [WARNING] Headers has beenwritten。想用 422 覆盖状态代码 400。如果您希望对行为有更大的控制，请考虑使用 ShouldBind 等效方法。
//类型 - 应该绑定 ShouldBind
//方法 - ShouldBind、ShoulBindJSON、ShouldBindXML、ShouldBindQuery、ShouldBindYAML、ShouldBindHeader
//行为 - 这些方法在幕后使用 ShouldBindWith。如果存在绑定错误，则返回错误，并且开发人员有责任适当地处理请求和错误。
//使用 Bind-method 时，Gin 尝试根据 Content-Type 标头推断绑定器。如果您确定要绑定的是什么，则可以使用 MustBindWith 或 ShouldBindWith。
//
//您还可以指定特定字段是必需的。如果某个字段使用 binding:"required" 修饰并且在绑定时为空值，则会返回错误。

require github.com/gin-gonic/gin v1.7.2
