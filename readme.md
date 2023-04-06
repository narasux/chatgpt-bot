# chatgpt-bot

## 机器人功能

🗣 语音交流：私人直接与机器人畅所欲言

💬 多话题对话：支持私人和群聊多话题讨论，高效连贯

🖼 文本成图：支持文本成图和以图搜图

🛖 场景预设：内置丰富场景列表，一键切换AI角色

🎭 角色扮演：支持场景模式，增添讨论乐趣和创意

🔄 上下文保留：回复对话框即可继续同一话题讨论

⏰ 自动结束：超时自动结束对话，支持清除讨论历史

📝 富文本卡片：支持富文本卡片回复，信息更丰富多彩

👍 交互式反馈：即时获取机器人处理结果

🎰 余额查询：即时获取 token 消耗情况

🔙 历史回档：轻松回档历史对话，继续话题讨论 🚧

🔒 管理员模式：内置管理员模式，使用更安全可靠 🚧

🌐 多token负载均衡：优化生产级别的高频调用场景

↩️ 支持反向代理：为不同地区的用户提供更快、更稳定的访问体验

📚 与飞书文档互动：成为企业员工的超级助手 🚧

🎥 话题内容秒转PPT：让你的汇报从此变得更加简单 🚧

📊 表格分析：轻松导入飞书表格，提升数据分析效率 🚧

🍊 私有数据训练：利用公司产品信息对GPT二次训练，更好地满足客户个性化需求 🚧

## 项目部署

### docker 部署

```bash
docker build -t feishu-chatgpt:latest .
docker run -d --name feishu-chatgpt -p 9000:9000 \
--env APP_ID=xxx \
--env APP_SECRET=xxx \
--env APP_ENCRYPT_KEY=xxx \
--env APP_VERIFICATION_TOKEN=xxx \
--env BOT_NAME=gpt \
--env OPENAI_KEY="sk-xxx1,sk-xxx2,sk-xxx3" \
--env API_URL="https://api.openai.com" \
--env HTTP_PROXY="" \
feishu-chatgpt:latest
```

注意:

- `BOT_NAME` 为飞书机器人名称，例如 `gpt`
- `OPENAI_KEY` 为 openai key，多个 key 用逗号分隔，例如 `sk-xxx1,sk-xxx2,sk-xxx3`
- `HTTP_PROXY` 为宿主机的 proxy 地址，例如 `http://host.docker.internal:7890`，没有代理的话，可以不用设置
- `API_URL` 为openai api 接口地址，例如 `https://api.openai.com`，没有反向代理的话，可以不用设置

### docker-compose 部署

编辑 docker-compose.yaml，通过 environment 配置相应环境变量（或者通过 volumes 挂载相应配置文件），然后运行下面的命令即可

```shell
# 构建镜像
docker compose build

# 启动服务
docker compose up -d

# 停止服务
docker compose down
```

## 详细配置步骤

- 获取 [OpenAI](https://platform.openai.com/account/api-keys) 的 KEY
- 创建 [飞书](https://open.feishu.cn/) 机器人
  - 前往[开发者平台](https://open.feishu.cn/app?lang=zh-CN)创建应用，并获取到 `APP_ID` 和 `APP_SECRET`
  - 前往 `应用功能-机器人`, 创建机器人并获取到 `ENCRYPT_KEY` 和 `VERIFICATION_TOKEN`
  - 在飞书机器人后台的 `事件订阅` 板块填写事件回调地址：`http://127.0.0.1/webhook/event`
  - 在飞书机器人后台的 `机器人` 板块填写消息卡片回调地址：`http://127.0.0.1/webhook/card`
  - 在事件订阅板块，搜索三个词 `机器人进群`、 `接收消息`、 `消息已读`, 把他们后面所有的权限全部勾选
  - 进入权限管理界面，搜索`图片`, 勾选`获取与上传图片或文件资源`。 
- 发布版本，申请上线，等待企业管理员审核通过
