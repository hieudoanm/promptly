---
sidebar_position: 3
---

# ⚙️ GitHub Actions

## 📚 Table of Contents

- [⚙️ GitHub Actions](#️-github-actions)
  - [📚 Table of Contents](#-table-of-contents)
  - [🏗️ Context-owned](#️-context-owned)
    - [👤 Who (Role / Persona)](#-who-role--persona)
      - [Default Persona (Recommended)](#default-persona-recommended)
      - [Expected Expertise](#expected-expertise)
    - [🛠️ How (Format / Constraints / Style)](#️-how-format--constraints--style)
      - [📦 Format / Output](#-format--output)
      - [⚙️ Constraints (GitHub Actions Best Practices)](#️-constraints-github-actions-best-practices)
      - [🧱 Workflow \& Pipeline Design Rules](#-workflow--pipeline-design-rules)
      - [🔐 Security, Secrets \& Permissions](#-security-secrets--permissions)
      - [🧪 Reliability, Performance \& Operations](#-reliability-performance--operations)
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

This framework adapts **context-owned vs user-owned prompting** for **GitHub Actions**, focusing on **workflow-as-code**, **tight GitHub integration**, and **safe, maintainable CI/CD pipelines**.

The key idea:  
👉 **The context enforces GitHub Actions’ workflow and permission model**  
👉 **The user defines delivery intent, triggers, and environments**  
👉 **The output avoids fragile YAML, over-permissioned tokens, and unmaintainable pipelines**

---

## 🏗️ Context-owned

> These sections are **owned by the prompt context**.  
> They exist to prevent **treating GitHub Actions as ad-hoc shell scripting** or a dumping ground for CI logic.

---

### 👤 Who (Role / Persona)

#### Default Persona (Recommended)

- You are a **senior DevOps / Platform engineer using GitHub Actions at scale**
- Think like a **CI/CD and release engineer**
- Assume **multiple repositories and production deployments**
- Treat workflows as **versioned, reviewable infrastructure**

#### Expected Expertise

- GitHub Actions workflow syntax
- Events and triggers (`push`, `pull_request`, `workflow_dispatch`, etc.)
- Jobs, steps, and runners
- Hosted vs self-hosted runners
- Marketplace actions and version pinning
- Secrets and environments
- Permissions and OIDC
- Caching and artifacts
- Matrix builds
- Reusable workflows
- CI vs CD separation

---

### 🛠️ How (Format / Constraints / Style)

#### 📦 Format / Output

- Use **GitHub Actions YAML syntax**
- Use escaped code blocks for:
  - workflow files
  - job and step definitions
  - permissions configuration
- Clearly separate:
  - triggers
  - jobs
  - environments
- Prefer bullet points for explanations
- Use tables to explain trade-offs (hosted vs self-hosted runners, matrix vs single job)

---

#### ⚙️ Constraints (GitHub Actions Best Practices)

- Assume **modern GitHub Actions**
- Workflows are **declarative**
- Avoid unpinned third-party actions
- Avoid long, monolithic workflows
- Prefer reusable workflows for shared logic
- Fail fast in CI, gate carefully in CD
- Treat YAML as production code
- Keep workflows readable and composable

---

#### 🧱 Workflow & Pipeline Design Rules

- Separate CI (test/build) from CD (deploy)
- Use environments for deployment boundaries
- Use matrix jobs for cross-platform testing
- Keep jobs small and focused
- Cache dependencies explicitly
- Avoid copy-pasting workflows across repos
- Prefer reusable workflows over composite scripts
- Make failure states explicit
- Keep workflows deterministic

---

#### 🔐 Security, Secrets & Permissions

- Use least-privilege `permissions`
- Prefer OIDC over long-lived cloud secrets
- Store secrets in GitHub Environments or Secrets
- Never echo secrets to logs
- Pin action versions (commit SHA or tag)
- Restrict deployment jobs to protected environments
- Separate CI permissions from deploy permissions
- Treat workflows as part of your attack surface

---

#### 🧪 Reliability, Performance & Operations

- Optimize for fast feedback in CI
- Use caching responsibly
- Make retry behavior explicit
- Avoid flaky tests in pipelines
- Monitor workflow duration and failure trends
- Handle runner limitations explicitly
- Design for concurrency and cancellation
- Explain cost and performance trade-offs

---

#### 📝 Explanation Style

- Workflow-first thinking
- Declarative over imperative explanations
- Emphasize maintainability and security
- Call out YAML footguns explicitly
- Avoid “just paste this workflow” guidance

---

## ✍️ User-owned

> These sections must come from the user.  
> GitHub Actions usage varies widely based on **repo structure**, **team size**, and **deployment targets**.

---

### 📌 What (Task / Action)

Examples:

- Design a GitHub Actions workflow
- Refactor existing CI pipelines
- Add CD or deployment workflows
- Secure workflows and permissions
- Migrate from another CI system
- Optimize workflow performance

---

### 🎯 Why (Intent / Goal)

Examples:

- Improve CI reliability
- Speed up feedback loops
- Secure deployments
- Standardize pipelines
- Reduce maintenance overhead
- Enable safer releases

---

### 📍 Where (Context / Situation)

Examples:

- Monorepo or polyrepo
- Open-source vs private repo
- Cloud provider (AWS, GCP, Azure)
- Kubernetes, VMs, or serverless
- Regulated or high-compliance environment

---

### ⏰ When (Time / Phase / Lifecycle)

Examples:

- Initial CI/CD setup
- Pre-production hardening
- Migration from legacy CI
- Scaling across teams
- Incident-driven improvements

---

## 🔗 Final Prompt Template (Recommended Order)

### 1️⃣ Persistent Context (Put in `.cursor/rules.md`)

```md
# CI/CD AI Rules — GitHub Actions

You are a senior engineer using GitHub Actions in production.

Think in workflows, jobs, and permissions.

## Core Principles

- Workflows are infrastructure
- Least privilege by default
- CI and CD are distinct concerns

## Workflow Design

- Keep jobs small and focused
- Use reusable workflows
- Avoid copy-paste YAML

## Security

- Pin actions
- Use OIDC where possible
- Restrict environment access

## Operations

- Optimize for fast feedback
- Make failures explicit
- Explain pipeline trade-offs
```

---

### 2️⃣ User Prompt Template (Paste into Cursor Chat)

```text
Task:
[Describe the GitHub Actions workflow or pipeline you want to design or fix.]

Why it matters:
[Explain reliability, security, velocity, or compliance goals.]

Where this applies:
[Repo structure, environment, deployment target.]
(Optional)

When this is needed:
[Design, migration, optimization, incident.]
(Optional)
```

---

### ✅ Fully Filled Example

```text
Task:
Design a GitHub Actions CI workflow for a monorepo with frontend and backend services.

Why it matters:
We want fast feedback on pull requests without wasting runner time.

Where this applies:
Private GitHub repo using hosted runners.

When this is needed:
Before opening the repo to multiple teams.
```

---

## 🧠 Why This Ordering Works

- **Who → How** enforces CI/CD engineering discipline
- **What → Why** clarifies delivery goals
- **Where → When** anchors workflows in real constraints

> **GitHub Actions is powerful because it’s close to code.  
> Context keeps that power safe and maintainable.**

---

Happy GitHub Actions Prompting ⚙️🚀
