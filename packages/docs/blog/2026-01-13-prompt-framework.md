---
slug: prompt-framework
title: Prompt Framework
authors: hieudoan
tags: [hello]
---

## 📚 Table of Contents

- [📚 Table of Contents](#-table-of-contents)
- [🧠✨ Introduction](#-introduction)
- [🧩 5W1H](#-5w1h)
  - [👤 **Who**](#-who)
  - [📌 **What**](#-what)
  - [🎯 **Why**](#-why)
  - [⏱️ **When** _(optional)_](#️-when-optional)
  - [🌍 **Where** _(optional)_](#-where-optional)
  - [⚙️ **How**](#️-how)
- [✅ Good Prompt vs ❌ Bad Prompt](#-good-prompt-vs--bad-prompt)
  - [❌ Bad Prompt](#-bad-prompt)
  - [✅ Good Prompt](#-good-prompt)
- [🔗 Combination](#-combination)
  - [🧠 Prompt Formula](#-prompt-formula)
  - [🔁 Reusable Prompt Template](#-reusable-prompt-template)
  - [4️⃣ Why This Works](#4️⃣-why-this-works)
- [Final Thoughts 🚀](#final-thoughts-)

## 🧠✨ Introduction

🧠 Prompting is quickly becoming a **core skill** for developers, designers, and knowledge workers.

⏳ A good prompt can save hours of work, while a vague one can lead to confusing or unusable results.

🧩 Think of prompts as **interfaces for thinking** — the clearer your interface, the better the output.

🚀 This post introduces a **simple, reusable Prompt Framework** that helps you to:

- 🗣️ communicate intent clearly
- 🎯 get more consistent results
- 🔁 reuse prompts across tools and projects

---

<!-- truncate -->

## 🧩 5W1H

🧠 One of the easiest ways to structure a strong prompt is by borrowing the classic **5W1H framework**.

### 👤 **Who**

👤 Who is the AI supposed to be?

- A senior engineer
- A product manager
- A teacher explaining to beginners

> _“Act as a senior backend engineer…”_

---

### 📌 **What**

📌 What exactly do you want?

- Generate code
- Review text
- Explain a concept

> _“Explain how Kafka consumer groups work…”_

---

### 🎯 **Why**

🎯 Why are you asking?

- Learning
- Production use
- Decision making

🧠 This helps the model choose the **right depth and tone**.

> _“…so I can explain it to junior developers.”_

---

### ⏱️ **When** _(optional)_

⏱️ When are time or context constraints relevant?

- Current trends
- Backward compatibility
- Version-specific behavior

> _“Using Kafka 3.x…”_

---

### 🌍 **Where** _(optional)_

🌍 Where will this be used?

- Web
- Mobile
- Backend
- React, Spring Boot, FastAPI

> _“…in a Spring Boot application.”_

---

### ⚙️ **How**

⚙️ How should the output look?

- Bullet points
- Code only
- Step-by-step explanation

> _“Use simple examples and diagrams in text.”_

---

## ✅ Good Prompt vs ❌ Bad Prompt

### ❌ Bad Prompt

> “Explain Kafka”

🚨 Problems:

- 👤 No role
- 🎯 No goal
- 🌍 No context
- ⚙️ No format
- ❌ No example

🤔 The AI must **guess everything**, which often leads to generic or unfocused answers.

---

### ✅ Good Prompt

> “Act as a senior backend engineer.  
> Explain Kafka consumer groups in simple terms for junior developers, using Kafka 3.x and a Spring Boot context.  
> Use bullet points and a short example, similar to explaining how HTTP load balancing works.”

✅ Why this works:

- 👤 Clear role
- 🎯 Clear goal
- 🌍 Clear context
- 📐 Clear format
- 🧩 Clear example

🧠 Examples act as **anchors** — they show the AI _what “good” looks like_, not just what to do.

---

## 🔗 Combination

🔗 The real power comes from **combining frameworks**.

### 🧠 Prompt Formula

```text
Role + Task + Context + Constraints + Output Format + (Optional) Example
```

---

### 🔁 Reusable Prompt Template

```text
Act as [ROLE].

Your task is to [WHAT].
This is for [WHY].
Context: [WHERE / WHEN].

Output requirements:

- Format: [FORMAT]
- Style: [STYLE]
- Length: [LENGTH]
- Example (optional): [REFERENCE OR SAMPLE]
```

🔁 You can store these templates and reuse them across:

- ChatGPT
- Claude
- GitHub Copilot
- Internal AI tools

---

### 4️⃣ Why This Works

🧠 This framework works because it **separates structure from intent**.

- 🧩 **Who + How → quality control (template-owned)**  
  🎛️ These define _how_ the AI should think and respond.  
  📈 They improve output quality regardless of topic.

- 🎯 **What + Why → intent (user-owned)**  
  🧠 These define _what matters_.  
  🙋 Only the user knows the real task and success criteria.

- 🌍 **Where / When → relevance (user-owned, optional)**  
  🧭 These ground the response in a real-world context, when needed.

🧩 Examples sit at the boundary:

- Optional
- User-provided
- Extremely powerful when clarity or style matters

⚖️ By keeping:

- 🏗️ structure in the **template**
- ✍️ meaning in the **user’s hands**

🚀 you avoid over-constraining prompts while still getting **consistent, high-quality results**.

> ✨ **Template sets clarity. User sets purpose. Examples set direction.**

---

## Final Thoughts 🚀

🧠 Prompting is not about “talking nicely to AI”.  
🎯 It’s about **thinking clearly and expressing intent**.

✅ A good prompt:

- ❄️ reduces ambiguity
- 📈 improves output quality
- ⏳ saves time

🚀 Start simple. Use 5W1H.  
🔁 Then combine and refine as you go.

> ✨ **Clear thinking → Clear prompts → Better results**

😄 Happy prompting!
