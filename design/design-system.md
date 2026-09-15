# Design System — Hello World Acceptance 7

> Source of truth: approved `index.html`.
> Every value below is extracted from it. Changing a value here without changing approved design is defect.

Last updated: 2026-09-15

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background, input background, button text |
| `--color-text` | `#000000` | Heading, form text, message text, input border |
| `--color-primary-action` | `#2563EB` | Save button background and border |
| `--color-focus` | `#2563EB` | Focus ring on input and button |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA |
| `--color-bg` | `--color-primary-action` | `5.17:1` | AA |
| `--color-primary-action` | `--color-bg` | `5.17:1` | AA for focus/UI |

### 1.2 Spacing

Base unit: `1px`. Every margin, padding, and gap in approved design uses one of these.

| Token | Value |
|---|---|
| `--space-0` | `0` |
| `--space-1` | `1px` |
| `--space-3` | `3px` |
| `--space-12` | `12px` |
| `--space-14` | `14px` |
| `--space-20` | `20px` |
| `--space-24` | `24px` |
| `--space-32` | `32px` |

### 1.3 Typography

Font families:

- Body: `Arial, Helvetica, sans-serif`, system-installed fonts.
- Headings: `Arial, Helvetica, sans-serif`, inherited from body.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-message` | `14px` | `1.5` inherited | `400` inherited | Polite form message |
| `--text-body` | Browser default `16px` | `1.5` | `400` inherited | Input and button text |
| `--text-heading` | `clamp(40px, 10vw, 72px)` | `1.05` | `700` | h1 greeting |

Heading levels are used in order: one `h1`, no skipped heading level.

| Token | Value | Used for |
|---|---|---|
| `--font-weight-body` | `400` inherited | Input and message text |
| `--font-weight-heading` | `700` | Greeting heading |
| `--font-weight-action` | `700` | Save button label |

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-square` | `0` | Input and button corners |
| `--border-width` | `1px` | Input and button border |
| `--focus-ring-width` | `3px` | Input and button focus outline |
| `--focus-ring-offset` | `3px` | Input and button focus outline offset |

Motion: approved design uses no animation or transitions.

### 1.5 Layout and breakpoints

| Name | Max width | Container | Columns | Gutter |
|---|---|---|---|---|
| `compact-form` | `520px` | `min(100%, 520px)` | Form changes from row to column | `12px` |

No z-index scale is used in approved design.

## 2. Components

### 2.1 Greeting Section

**Purpose** — Center one editable persisted greeting. Do not use for multi-section pages or navigation.

**Anatomy** — `[heading] [form] [polite message]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default centered | `--color-bg`, `--color-text`, `--space-24`, `--space-32`, `--text-heading` | Only page section |

**Sizes**

| Size | Container | Padding | Text token |
|---|---|---|---|
| Default | `min(100%, 520px)` | `24px` on `main` | `--text-heading` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Large centered heading, form below, message area reserved | `--color-bg`, `--color-text`, `--space-24`, `--space-32` |

**Accessibility** — Section uses `aria-labelledby` pointing to `h1`. Message uses `aria-live="polite"`.

### 2.2 Text Input

**Purpose** — Edit greeting text.

**Anatomy** — `[visually hidden label] [text input]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default text | `--color-bg`, `--color-text`, `--border-width`, `--radius-square` | Greeting entry |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | Content-driven | `12px 14px` | `--text-body` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White background, black text, black 1px border | `--color-bg`, `--color-text`, `--border-width`, `--radius-square` |
| Focus | Blue 3px outline with 3px offset | `--color-focus`, `--focus-ring-width`, `--focus-ring-offset` |

**Accessibility** — Native text input with hidden visible label text `Greeting`; `required`; keyboard focus uses `:focus-visible`; input is at least 44px tall from padding plus line height.

### 2.3 Primary Button

**Purpose** — Save current greeting. Use for primary page action only.

**Anatomy** — `[label]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Primary action | `--color-primary-action`, `--color-bg`, `--border-width`, `--radius-square`, `--font-weight-action` | Save action |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | Content-driven | `12px 20px` | `--text-body` |
| Compact viewport | Full form width | `12px 20px` | `--text-body` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Blue background, blue border, white bold text, pointer cursor | `--color-primary-action`, `--color-bg`, `--border-width`, `--font-weight-action` |
| Focus | Blue 3px outline with 3px offset | `--color-focus`, `--focus-ring-width`, `--focus-ring-offset` |

**Accessibility** — Native submit button; keyboard focus uses `:focus-visible`; minimum hit target is at least 44px tall from padding plus line height.

### 2.4 Form Message

**Purpose** — Show save confirmation or validation message below form.

**Anatomy** — `[message text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Polite status | `--color-text`, `--text-message`, `--space-12`, `--space-24` | Feedback after submit |

**Sizes**

| Size | Min height | Margin | Text token |
|---|---|---|---|
| Default | `24px` | `12px 0 0` | `--text-message` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Empty | Reserved 24px message area with no visible text | `--space-24` |
| Saved | Text reads `Saved.` | `--color-text`, `--text-message` |
| Validation error | Text reads `Enter a greeting before saving.` | `--color-text`, `--text-message` |

**Accessibility** — `aria-live="polite"` announces message changes without moving focus, except validation error returns focus to input.

## 3. Content and formatting

- Voice and tone: plain, direct, minimal.
- Locale: English UI copy; no date, time, number, or currency formats used.
- Capitalization: heading preserves greeting text exactly; button uses title case as approved: `Save`; label uses title case: `Greeting`.
- Empty-state and error-message wording: short sentence stating required action, e.g. `Enter a greeting before saving.`

## 4. Known deviations

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Spacing scale | Uses exact values `1px`, `3px`, `14px`, and `20px` rather than common 4px scale | Approved mockup defines these values for borders, focus ring, input padding, and button padding | Keep for this app unless design is revised |
| Button and input states | Approved design draws default and focus states only; no hover, active, or disabled visual states | Documenting extra states would invent scope | Add states only if approved design changes |
| AI defaults check | Plain blue action color is close to generic AI defaults, but no purple/indigo palette, gradients, heavy shadows, emoji, filler copy, or generic multi-section layout are used | Stakeholder requested minimal white/black design with blue `#2563EB` button | No follow-up |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2026-09-15 | Initial design system extracted from approved `index.html` | This PR |
