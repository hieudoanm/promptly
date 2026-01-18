# [Promptly](https://hieudoanm.github.io/promptly)

- [Promptly](#promptly)
  - [Demo](#demo)
  - [Rules](#rules)
  - [Models](#models)

## Demo

[Promptly - Chat](https://hieudoanm-promptly.vercel.app)

## Rules

Total of **142** Cursor Rules for Indie/Solo Developers.

```shell
|-- engineering
|  |-- data
|  |  |-- analyst
|  |  |  |-- non-technical
|  |  |  |  |-- powerbi.md
|  |  |  |  |-- tableau.md
|  |  |  |-- technical
|  |  |  |  |-- matplotlib.md
|  |  |  |  |-- numpy.md
|  |  |  |  |-- pandas.md
|  |  |  |  |-- statsmodels.md
|  |  |-- engineer
|  |  |  |-- apache-airflow.md
|  |  |  |-- apache-iceberg.md
|  |  |  |-- apache-spark.md
|  |  |  |-- apache-trino.md
|  |  |-- scientist
|  |  |  |-- hugging-face.md
|  |  |  |-- pytorch.md
|  |  |  |-- scikit-learn.md
|  |  |  |-- tensorflow.md
|  |  |  |-- xgboost.md
|  |-- developer-tools
|  |  |-- api-clients
|  |  |  |-- bruno.md
|  |  |  |-- insomnia.md
|  |  |  |-- postman.md
|  |  |-- code-editors
|  |  |  |-- cursor.md
|  |  |  |-- vscode.md
|  |  |  |-- vscodium.md
|  |  |  |-- windsurf.md
|  |  |-- ide
|  |  |  |-- android-studio.md
|  |  |  |-- jetbrains
|  |  |  |  |-- clion.md
|  |  |  |  |-- intellij-idea.md
|  |  |  |  |-- php-storm.md
|  |  |  |  |-- py-charm.md
|  |  |  |  |-- rider.md
|  |  |  |  |-- ruby-mine.md
|  |  |  |  |-- rust-rover.md
|  |  |  |  |-- web-storm.md
|  |  |  |-- visual-studio.md
|  |  |  |-- xcode.md
|  |-- devops
|  |  |-- ci-cd
|  |  |  |-- harness.md
|  |  |  |-- jenkins.md
|  |  |-- container
|  |  |  |-- docker.md
|  |  |  |-- kubernetes.md
|  |  |-- iac
|  |  |  |-- aws-cloudformation.md
|  |  |  |-- terraform.md
|  |  |-- observability
|  |  |  |-- aws-cloudwatch.md
|  |  |  |-- datadog.md
|  |  |  |-- grafana.md
|  |  |  |-- splunk.md
|  |  |-- secrets
|  |  |  |-- aws-secrets-manager.md
|  |  |  |-- azure-key-vault.md
|  |  |  |-- hashicorp-vault.md
|  |  |  |-- kubernetes-secrets.md
|  |-- game
|  |  |-- engines
|  |  |  |-- cocos.md
|  |  |  |-- godot.md
|  |  |  |-- unity.md
|  |  |  |-- unreal.md
|  |-- languages
|  |  |-- c.md
|  |  |-- cplusplus.md
|  |  |-- csharp.md
|  |  |-- go.md
|  |  |-- java.md
|  |  |-- kotlin.md
|  |  |-- php.md
|  |  |-- python.md
|  |  |-- ruby.md
|  |  |-- rust.md
|  |  |-- scala.md
|  |  |-- swift.md
|  |  |-- typescript.md
|  |-- software
|  |  |-- backend
|  |  |  |-- architecture
|  |  |  |  |-- cqrs.md
|  |  |  |  |-- event-driven.md
|  |  |  |  |-- hexagonal.md
|  |  |  |  |-- microservices.md
|  |  |  |  |-- monolith.md
|  |  |  |-- database
|  |  |  |  |-- cache
|  |  |  |  |  |-- memcached.md
|  |  |  |  |  |-- redis.md
|  |  |  |  |-- nosql
|  |  |  |  |  |-- dynamodb.md
|  |  |  |  |  |-- mongodb.md
|  |  |  |  |-- search
|  |  |  |  |  |-- elasticsearch.md
|  |  |  |  |  |-- opensearch.md
|  |  |  |  |-- sql
|  |  |  |  |  |-- postgresql.md
|  |  |  |  |  |-- sqlite.md
|  |  |  |-- events
|  |  |  |  |-- apache-kafka.md
|  |  |  |  |-- apache-pulsar.md
|  |  |  |  |-- rabbitmq.md
|  |  |  |-- languages
|  |  |  |  |-- csharp
|  |  |  |  |  |-- dotnet.md
|  |  |  |  |-- go
|  |  |  |  |  |-- gin.md
|  |  |  |  |-- jvm
|  |  |  |  |  |-- java
|  |  |  |  |  |  |-- spring-boot.md
|  |  |  |  |  |-- kotlin
|  |  |  |  |  |  |-- ktor.md
|  |  |  |  |  |-- scala
|  |  |  |  |  |  |-- play.md
|  |  |  |  |-- php
|  |  |  |  |  |-- laravel.md
|  |  |  |  |-- python
|  |  |  |  |  |-- fastapi.md
|  |  |  |  |-- ruby
|  |  |  |  |  |-- rails.md
|  |  |  |  |-- typescript
|  |  |  |  |  |-- express.md
|  |  |  |-- security
|  |  |  |  |-- jwt.md
|  |  |  |  |-- oauth2.md
|  |  |  |  |-- oidc.md
|  |  |-- cli
|  |  |  |-- cobra.md
|  |  |  |-- oclif.md
|  |  |-- frontend
|  |  |  |-- bff
|  |  |  |  |-- graphql.md
|  |  |  |  |-- trpc.md
|  |  |  |-- hybrid
|  |  |  |  |-- flutter.md
|  |  |  |  |-- ionic.md
|  |  |  |  |-- react-native.md
|  |  |  |  |-- tauri.md
|  |  |  |-- native
|  |  |  |  |-- desktop
|  |  |  |  |  |-- arch-linux.md
|  |  |  |  |  |-- macos.md
|  |  |  |  |  |-- ubuntu.md
|  |  |  |  |  |-- windows.md
|  |  |  |  |-- mobile
|  |  |  |  |  |-- android.md
|  |  |  |  |  |-- ios.md
|  |  |  |  |  |-- kaios.md
|  |  |  |  |  |-- ubuntu-touch.md
|  |  |  |-- web
|  |  |  |  |-- framework
|  |  |  |  |  |-- csr
|  |  |  |  |  |  |-- angular.md
|  |  |  |  |  |  |-- react.md
|  |  |  |  |  |  |-- solid.md
|  |  |  |  |  |  |-- svelte.md
|  |  |  |  |  |  |-- vue.md
|  |  |  |  |  |-- ssr
|  |  |  |  |  |  |-- angular-ssr.md
|  |  |  |  |  |  |-- next.md
|  |  |  |  |  |  |-- nuxt.md
|  |  |  |  |  |  |-- solid-start.md
|  |  |  |  |  |  |-- svelte-kit.md
|  |  |  |  |-- styling
|  |  |  |  |  |-- css-in-js
|  |  |  |  |  |  |-- emotion.md
|  |  |  |  |  |  |-- styled-components.md
|  |  |  |  |  |-- css-utilities
|  |  |  |  |  |  |-- tailwindcss.md
|  |  |  |  |  |  |-- unocss.md
|  |  |  |  |-- testing
|  |  |  |  |  |-- e2e
|  |  |  |  |  |  |-- cypress.md
|  |  |  |  |  |  |-- playwright.md
|  |  |  |  |  |-- unit
|  |  |  |  |  |  |-- jest.md
|  |  |  |  |  |  |-- vitest.md
|  |-- solution
|  |  |-- architect.md
|  |  |-- design.md
|-- intro.md
|-- marketing
|  |-- copy-writer
|  |  |-- sites
|  |  |  |-- acquire.md
|  |  |  |-- hacker-news.md
|  |  |  |-- indie-hackers.md
|  |  |  |-- product-hunt.md
|  |  |-- social
|  |  |  |-- github.md
|  |  |  |-- linkedin.md
|  |  |  |-- reddit.md
|  |  |  |-- twitter.md
|  |-- ui-ux-designer
|  |  |-- canva.md
|  |  |-- figma.md
|-- product
|  |-- business
|  |  |-- business-analyst.md
|  |  |-- product-owner.md
|  |-- delivery
|  |  |-- release-train-engineer.md
|  |  |-- scrum-master.md
```

## Models

| No  | Type  | Name                                      | Maintainer                           |
| --- | ----- | ----------------------------------------- | ------------------------------------ |
| 01  | Image | [Midjourney](https://www.midjourney.com)  |                                      |
| 02  | Text  | [Perplexity](https://www.perplexity.ai/)  |                                      |
| 03  | Text  | [Grok](https://grok.com/)                 | [X.Ai](https://x.ai/)                |
| 04  | Text  | [Poe](https://www.poe.com/)               | [Quora](https://www.quora.com/)      |
| 05  | Text  | [Claude](https://claude.ai/)              | [Anthropic](https://anthropic.com/)  |
| 06  | Text  | [Deepseek](https://chat.deepseek.com/)    | [DeepSeek AI](https://deepseek.ai/)  |
| 07  | Text  | [Gemini](https://gemini.google.com/)      | [Google](https://google.com/)        |
| 08  | Text  | [Copilot](https://copilot.microsoft.com/) | [Microsoft](https://microsoft.com/)  |
| 09  | Text  | [ChatGPT](https://chatgpt.com/)           | [OpenAI](https://openai.com/)        |
| 10  | Text  | [Llama](https://www.llama.com/)           | [Meta](https://developers.meta.com/) |
