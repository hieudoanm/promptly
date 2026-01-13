---
slug: prompt-framework
title: Prompt Framework
authors: hieudoan
tags: [hello]
---

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Introduction 🧠✨](#introduction-)
- [5W1H 🧩](#5w1h-)
  - [**Who** 👤](#who-)
  - [**What** 📌](#what-)
  - [**Why** 🎯](#why-)
  - [**When** ⏱️ *(optional)*](#when-️-optional)
  - [**Where** 🌍 *(optional)*](#where--optional)
  - [**How** ⚙️](#how-️)
- [Good Prompt ✅ vs Bad Prompt ❌](#good-prompt--vs-bad-prompt-)
  - [❌ Bad Prompt](#-bad-prompt)
  - [✅ Good Prompt](#-good-prompt)
- [Combination 🔗](#combination-)
  - [🧠 Prompt Formula](#-prompt-formula)
  - [🔁 Reusable Prompt Template](#-reusable-prompt-template)
- [Final Thoughts 🚀](#final-thoughts-)

## Introduction 🧠✨

Prompting is quickly becoming a **core skill** for developers, designers, and knowledge workers.  
A good prompt can save hours of work, while a vague one can lead to confusing or unusable results.

Think of prompts as **interfaces for thinking** — the clearer your interface, the better the output.

This post introduces a **simple, reusable Prompt Framework** that helps you to:

- communicate intent clearly 🗣️
- get more consistent results 🎯
- reuse prompts across tools and projects 🔁

---

## 5W1H 🧩

One of the easiest ways to structure a strong prompt is by borrowing the classic **5W1H framework**.

### **Who** 👤  

Who is the AI supposed to be?

- A senior engineer
- A product manager
- A teacher explaining to beginners

> *“Act as a senior backend engineer…”*

---

### **What** 📌

What exactly do you want?

- Generate code
- Review text
- Explain a concept

> *“Explain how Kafka consumer groups work…”*

---

### **Why** 🎯

Why are you asking?

- Learning
- Production use
- Decision making

This helps the model choose the **right depth and tone**.

> *“…so I can explain it to junior developers.”*

---

### **When** ⏱️ *(optional)*

Time or context constraints.

- Current trends
- Backward compatibility
- Version-specific behavior

> *“Using Kafka 3.x…”*

---

### **Where** 🌍 *(optional)*

Environment or platform.

- Web
- Mobile
- Backend
- React, Spring Boot, FastAPI

> *“…in a Spring Boot application.”*

---

### **How** ⚙️

How should the output look?

- Bullet points
- Code only
- Step-by-step explanation

> *“Use simple examples and diagrams in text.”*

---

## Good Prompt ✅ vs Bad Prompt ❌

### ❌ Bad Prompt

> “Explain Kafka”

Too broad. Too vague. Too much guessing.

---

### ✅ Good Prompt

> “Act as a senior backend engineer.  
> Explain Kafka consumer groups in simple terms for junior developers, using Kafka 3.x and a Spring Boot context.  
> Use bullet points and a short example.”

Clear role 👤  
Clear goal 🎯  
Clear format 📐  

---

## Combination 🔗

The real power comes from **combining frameworks**.

### 🧠 Prompt Formula

```text
Role + Task + Context + Constraints + Output Format
```

---

### 🔁 Reusable Prompt Template

```text
Act as [ROLE].

Your task is to [WHAT].  
This is for [WHY].  
Context: [WHERE / WHEN].

Output requirements:
- [FORMAT]
- [STYLE]
- [LENGTH]
```

You can store these templates and reuse them across:

- ChatGPT
- Claude
- GitHub Copilot
- Internal AI tools

---

## Final Thoughts 🚀

Prompting is not about “talking nicely to AI”.  
It’s about **thinking clearly and expressing intent**.

A good prompt:

- reduces ambiguity ❄️
- improves output quality 📈
- saves time ⏳

Start simple. Use 5W1H.  
Then combine and refine as you go.

> **Clear thinking → Clear prompts → Better results** ✨

---

Happy prompting! 😄
