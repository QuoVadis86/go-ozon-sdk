# 收到新版 Swagger 后如何维护 SDK

本文件是给下一位维护者或 AI 的完整交接说明。输入可以只有一份新版 `swagger.json`；不能假设维护者知道以前的 issue 或聊天记录。`AGENTS.md` 是入口，本文件保存具体操作和历史原因。

## 先认识仓库

- `origin/swagger.json` 是提交到仓库的 Ozon Seller Swagger 快照。`gen.go` 从这里生成各服务目录的 `types.go`、`service.go` 和 `service_test.go`。
- `product/attributes_v4_json.go`、`product/attributes_v3_compat.go`、`returns/legacy_rfbs.go` 是**非生成文件**。生成器不应覆盖它们。
- `client.go` 暴露 SDK 入口，`transport/` 处理 HTTP 和语言请求头。公开的 `NewClient(clientID, apiKey, opts)` 调用方式必须保持兼容。
- 生成器把已有方法、类型、字段保留在原顺序，新内容按名称排序后追加。不要重新全量按字母排序，避免每次更新出现数千行无意义重排。

## 1. 保存并核验输入

从用户提供的新版 Swagger 获取 JSON，保存到 `origin/swagger.json`。如果下载需要登录，仅把 JSON 文件提交到仓库；不要把 `Cookie`、访问令牌、请求头或带签名的 URL 写入代码、文档、日志或提交。先保留旧快照：

```sh
git show HEAD:origin/swagger.json > /tmp/ozon-swagger-before.json
# 将新版 JSON 放入 origin/swagger.json
python3 -m json.tool origin/swagger.json > /dev/null
python3 scripts/swagger_diff.py /tmp/ozon-swagger-before.json origin/swagger.json
```

报告会列出新增/移除路径、相同 `operationId` 的路径迁移、新增废弃标记、标签与请求/响应模型引用变化。逐项判断：路径拼写修正不等于新接口；标签变化可能把方法从 `beta` 移到 `fbs`；从 Swagger 消失不等于已经实测线上下线。必要时对照 Ozon 官方文档或经授权的真实响应，记录证据。

## 2. 更新生成代码，但保护公开 API

运行 `go run gen.go`。随后检查生成变动，尤其是 `service.go` 删除的方法、公开方法的参数/返回类型、`types.go` 删除或改名的类型及字段：

```sh
git diff -- '*/service.go' | rg '^[+-]func '
git diff -- '*/types.go' | rg '^[+-]type '
```

不要以“能通过仓库测试”为唯一标准：仓库内部可能没有使用旧公开类型，外部用户仍会编译失败。如果命令没有匹配项，`rg` 返回退出码 1，表示没有这种差异。

对于 Swagger 删除的旧路径，优先把旧方法和必要请求/响应类型留在独立的非生成文件，并添加独立段落的 Go 文档标记：

```go
// OldMethod explains the old API.
//
// Deprecated: this endpoint is absent from the current Ozon Swagger; use the documented replacement.
```

如果没有确定的新接口，不要编造替代方案；说明“已从当前 Swagger 移除，线上可用性未经验证”。`product/attributes_v3_compat.go` 保留旧的商品特征方法并转发 V4；`returns/legacy_rfbs.go` 保留五条已移除的 rFBS 方法，但仍请求旧路径。标签迁移也要检查旧 SDK 包中的方法是否需要兼容入口。只有明确决定进行破坏性版本升级时，才删除这些兼容代码。

生成器会根据 Swagger 的 `deprecated` 标记以及描述中的“停用/弃用/废弃”输出提示。`Deprecated:` 必须是注释中的**独立段落**，否则编辑器可能不识别；`deprecation_test.go` 检查当前废弃方法。新版如果出现其他废弃措辞，要扩展识别规则并补测试。

## 3. 记住 Swagger 曾经写错的三种响应形状

`POST /v4/product/info/attributes` 曾在真实响应中返回：

| 字段 | Swagger 快照 | 已观察到的 JSON | SDK 兼容方式 |
| --- | --- | --- | --- |
| `images[]` | `type: "array of strings"`，但 `items` 引用对象模型 | URL 字符串或对象 | `GetProductAttributesResponseImage.UnmarshalJSON` 接收两种形状；URL 保存在 `.URL` |
| `sku` | `string` | 数字或字符串 | 保持公开 `SKU string`，解码时转换数字 |
| `total` | `string`、`format: int64` | 数字或字符串 | 保持公开 `Total string`，解码时转换数字 |

这些问题来自文档声明与真实响应不一致，单纯重新生成会重现错误。`product/attributes_v4_json.go` 负责三处解码，`gen.go` 为图片模型保留 `URL` 字段；`product/attributes_v4_test.go` 覆盖字符串、数字和两种图片形状。不要把兼容解码写入生成文件，也不要因新版 Swagger 仍写 `string` 就删除它。只有拿到可信的新响应证据且验证兼容性后才能改。

还要核对 `gen.go` 中的 `stableSchemaNames`：以前 Swagger 键的遍历顺序曾使 `seller.StatusEnum`、`warehouse.FirstMileType` 等公开类型被重命名或覆盖。`TopSku` 的实际请求路径曾需修正为 `/v1/question/top-sku`。这些特殊处理不可仅凭一次生成差异删除。

**当前已知兼容风险：**2026 年 11 月快照把五个商品实例操作的标签从 `BetaMethod` 改到了 FBS，生成方法因此从 `beta.Service` 移入 `fbs.Service`；旧 `beta.Service` 入口在 v0.3.2 尚未补齐。后续维护者应在发布下一版前评估旧包的兼容包装及相关类型别名，不能把“接口路径没变”误判为“SDK 公共 API 没变”。

## 4. 语言行为不能被生成更新破坏

语言头是 `x-o3-language`。公开的 `LanguageDefault`、`LanguageRU`、`LanguageEN`、`LanguageTR`、`LanguageZHHans` 是受限常量，`LanguageZHHans` 映射到准确值 `zh-Hans`。同一个 client 应通过 `WithLanguage(ctx, ...)` 按请求覆盖；原 `ctx` 不受影响。客户端默认值是可选项，不得修改 `NewClient` 既有签名。语言头可本地化部分属性名，但不能承诺所有错误正文都是中文。测试见 `client_language_test.go` 和 `transport/client_test.go`。

## 5. 验证与发布清单

```sh
go test gen.go gen_order_test.go
env -u OZON_CLIENT_ID -u OZON_API_KEY go test ./...
git diff --check
```

生成的 `service_test.go` 在设置了 `OZON_CLIENT_ID` 和 `OZON_API_KEY` 时可能访问真实接口；普通回归测试务必先移除这两个变量。按需单独运行 `product/attributes_v4_test.go`、`returns/legacy_rfbs_test.go` 和根目录的 `deprecation_test.go`。检查同一份 Swagger 重复运行 `go run gen.go` 后，生成文件字节不再变化。核对 `README.md`、`README_CN.md` 中的接口/模型数与服务表。

可将第一次生成后的跟踪文件差异保存到临时文件，再运行一次生成器并比较。这个检查会覆盖 `types.go`、`service.go` 和 `service_test.go`：

```sh
git diff -- . ':!origin/swagger.json' > /tmp/ozon-first-generation.diff
go run gen.go
git diff -- . ':!origin/swagger.json' > /tmp/ozon-second-generation.diff
cmp /tmp/ozon-first-generation.diff /tmp/ozon-second-generation.diff
```

发布前逐项回答：新增路径是否有方法；消失/改名的旧公开方法是否保留或明确说明；废弃方法是否有独立 `Deprecated:` 段落；公开 client 初始化方式和语言按请求切换是否仍兼容；V4 商品特征三种响应形状是否仍能解码。通过 PR 合并后，版本 tag 指向合并提交。不要把用户已有的本地改动混入提交。
