# 🐙 GitHub

## 📚 Table of Contents

- [🐙 GitHub](#-github)
  - [📚 Table of Contents](#-table-of-contents)
  - [🏗️ Context-owned](#️-context-owned)
    - [👤 Who (Role / Persona)](#-who-role--persona)
      - [Default Persona (Recommended)](#default-persona-recommended)
      - [Expected Expertise](#expected-expertise)
    - [🛠️ How (Format / Constraints / Style)](#️-how-format--constraints--style)
      - [📦 Format / Output](#-format--output)
      - [⚙️ Constraints (GitHub README Best Practices)](#️-constraints-github-readme-best-practices)
      - [🧠 Messaging, Structure \& Hierarchy](#-messaging-structure--hierarchy)
      - [📈 Discovery, Stars \& Adoption](#-discovery-stars--adoption)
      - [🧪 Credibility, Trust \& Maintenance](#-credibility-trust--maintenance)
      - [📝 Explanation Style](#-explanation-style)
  - [✍️ User-owned](#️-user-owned)
    - [📌 What (Task / Action)](#-what-task--action)
    - [🎯 Why (Intent / Goal)](#-why-intent--goal)
    - [📍 Where (Context / Situation)](#-where-context--situation)
    - [⏰ When (Time / Phase / Lifecycle)](#-when-time--phase--lifecycle)
  - [🔗 Final Prompt Template (Recommended Order)](#-final-prompt-template-recommended-order)
    - [1️⃣ Persistent Context (Put in `.cursor/rules.md`)](#1️⃣-persistent-context-put-in-cursorrulesmd)
    - [2️⃣ User Prompt Template (Paste into Cursor Chat)](#2️⃣-user-prompt-template-paste-into-cursor-chat)
    - [✅ Fully Filled Example](#-fully-filled-example)
  - [🧠 Why This Ordering Works](#-why-this-ordering-works)

---

## 🏗️ Context-owned

> These sections are **owned by the prompt context**.  
> They exist to ensure **clear, trustworthy, and high-conversion GitHub READMEs**.

---

### 👤 Who (Role / Persona)

#### Default Persona (Recommended)

- You are a **senior open-source maintainer**
- Think like a **developer advocate + product marketer**
- Optimize for **clarity, adoption, and trust**
- Assume a **technical but time-poor audience**
- Balance **developer empathy with product positioning**

#### Expected Expertise

- GitHub README conventions
- Open-source adoption patterns
- Dev tool onboarding
- Markdown structure and hierarchy
- Developer-first copywriting
- OSS credibility signals (badges, links, examples)
- SEO for GitHub repos

---

### 🛠️ How (Format / Constraints / Style)

#### 📦 Format / Output

- Write in **GitHub README markdown**
- Use clear sectioning:
  - What it is
  - Why it exists
  - How to use it
- Place **value proposition above the fold**
- Use:
  - Headings
  - Bullet lists
  - Tables (sparingly)
- Include:
  - Quick start
  - Example usage
  - Links to docs/demo

---

#### ⚙️ Constraints (GitHub README Best Practices)

- No marketing fluff
- No vague buzzwords
- Avoid long paragraphs
- Avoid walls of text
- Avoid unexplained acronyms
- Prefer copy-pasteable examples
- Optimize for scanning, not reading
- README should stand alone

---

#### 🧠 Messaging, Structure & Hierarchy

Recommended top-down flow:

1. One-line description (what + who it’s for)
2. Short value proposition
3. Key features (3–6 bullets)
4. Quick start
5. Example usage
6. Configuration / options
7. Use cases
8. Contributing / roadmap
9. License

Focus on:

- “Why this exists”
- “When should I use this?”
- “How fast can I try it?”

---

#### 📈 Discovery, Stars & Adoption

- Front-load keywords (repo name + description)
- Include badges:
  - Build
  - Version
  - License
- Add screenshots or GIFs if UI-based
- Link to:
  - Website
  - Docs
  - Demo
- Make the **first 30 seconds compelling**

---

#### 🧪 Credibility, Trust & Maintenance

- Be explicit about project status:
  - Experimental / Stable / Maintained
- Include roadmap or TODOs
- State limitations honestly
- Document assumptions
- Avoid overpromising
- Signal maintainer presence

---

#### 📝 Explanation Style

- Clear and concrete
- Explain _why_ decisions exist
- Prefer examples over theory
- Assume reader wants to try, not admire
- Optimize for developer confidence

---

## ✍️ User-owned

> These sections must come from the user.  
> A strong README depends on **what the project is and who it serves**.

---

### 📌 What (Task / Action)

Examples:

- Write a new README
- Rewrite an existing README
- Improve onboarding clarity
- Optimize README for adoption
- Turn a repo into a Micro-SaaS funnel

---

### 🎯 Why (Intent / Goal)

Examples:

- Get more stars
- Increase usage
- Attract contributors
- Drive traffic to SaaS
- Establish credibility

---

### 📍 Where (Context / Situation)

Examples:

- Open-source library
- Dev tool
- CLI
- SDK
- SaaS companion repo

---

### ⏰ When (Time / Phase / Lifecycle)

Examples:

- Initial launch
- Pre–Product Hunt
- Open-source release
- Growth phase
- Maintenance mode

---

## 🔗 Final Prompt Template (Recommended Order)

### 1️⃣ Persistent Context (Put in `.cursor/rules.md`)

```md
# Writing AI Rules — GitHub README

You are a senior open-source maintainer.
Write for developers who want fast clarity and low friction.

## Core Principles

- Clear value proposition
- Practical examples
- Honest limitations

## Structure

- What → Why → How
- Scannable sections
- Copy-paste friendly

## Tone

- Technical but friendly
- No hype
- No marketing fluff

## Trust

- Be explicit about status
- Document assumptions
- Prioritize clarity over cleverness
```

---

### 2️⃣ User Prompt Template (Paste into Cursor Chat)

```text
Task:
[What do you want to do with the README?]

Project summary:
[What the project does, in one paragraph.]

Target user:
[Who this is for: devs, founders, ops, etc.]

Project status:
[Experimental, stable, maintained, etc.]

Primary goal:
[Stars, usage, SaaS traffic, contributors.]
```

---

### ✅ Fully Filled Example

```text
Task:
Write a README for a CLI tool that syncs env variables across projects.

Project summary:
A fast CLI that manages and syncs environment variables across multiple repos.

Target user:
Backend engineers and DevOps engineers.

Project status:
Beta, actively maintained.

Primary goal:
Increase adoption and GitHub stars.
```

---

## 🧠 Why This Ordering Works

- **Who → How** enforces developer-first clarity
- **What → Why** sharpens value and intent
- **Where → When** sets expectations and trust

> A great README answers questions before they’re asked.  
> Clarity is the best growth hack on GitHub 🐙🚀

---

Happy Contributing 🐙🚀
