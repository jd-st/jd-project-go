# Shared Params Types

- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared#Order">Order</a>

# Pets

Params Types:

- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#CategoryParam">CategoryParam</a>
- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetParam">PetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#Category">Category</a>
- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#Pet">Pet</a>
- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetUploadImageResponse">PetUploadImageResponse</a>

Methods:

- <code title="post /pet">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetNewParams">PetNewParams</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetUpdateParams">PetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetFindByStatusParams">PetFindByStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetFindByTagsParams">PetFindByTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.UpdateByID">UpdateByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetUpdateByIDParams">PetUpdateByIDParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pets.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, image <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetUploadImageParams">PetUploadImageParams</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#PetUploadImageResponse">PetUploadImageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# St0re

Response Types:

- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#St0reListInventoryResponse">St0reListInventoryResponse</a>

Methods:

- <code title="get /st0re/inventory">client.St0re.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#St0reService.ListInventory">ListInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#St0reListInventoryResponse">St0reListInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

Methods:

- <code title="post /st0re/order">client.St0re.Orders.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#St0reOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#St0reOrderNewParams">St0reOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /st0re/order/{orderId}">client.St0re.Orders.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#St0reOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /st0re/order/{orderId}">client.St0re.Orders.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#St0reOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#User">User</a>

Methods:

- <code title="post /user">client.Users.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, existingUsername <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.Users.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserNewWithListParams">UserNewWithListParams</a>) (<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.Users.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jd-st/jd-project-go">jdproject</a>.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.Users.<a href="https://pkg.go.dev/github.com/jd-st/jd-project-go#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
