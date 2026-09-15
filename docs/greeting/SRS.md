# SRS — Greeting

Module: `greeting`
Design: [View the approved design](http://localhost:8080/design/2fc37493-97db-46a1-b75d-7f5246b02495)
Design system: `design/design-system.md`

> One file per module, at `docs/{module}/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

Greeting module lets any visitor view and edit one persisted greeting for "Hello World Acceptance 7". Without it, project cannot prove database persistence, Go API serving/updating, and Next.js rendering through end-to-end pipeline.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Any person using public page; no sign-in exists | View current greeting, edit greeting text, save non-empty greeting |

## 3. Scope

**In scope** — functions specified below, by plan title:

- Persisted Editable Greeting

**Out of scope** — expected adjacent work not built in this module:

- Sign-in and permissions — deliberately not built; brief says no sign-in.
- Navigation or extra sections — deliberately not built; approved design has one centered section only.
- External services — deliberately not built; brief says no external services.
- Multiple greetings or greeting history — deliberately not built; plan covers one stored greeting.

## 4. Functional requirements

### 4.1 Persisted Editable Greeting

**Requirement GREETING-001 — Show stored greeting**

*As a* Visitor, *I want to* see current greeting when page loads, *so that* page reflects stored product state.

Behaviour:

1. Visitor opens greeting page.
2. Page shows one centered section.
3. Large heading shows current stored greeting.
4. Text input value matches current stored greeting.
5. When no prior visitor-saved value exists, current stored greeting is `Hello, World!`.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through story plan criteria that cite it.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Stored greeting has initial value `Hello, World!` | Visitor opens page | Large heading text is `Hello, World!` |
| AC-2 | Stored greeting is `Hello, World!` | Visitor opens page | Text input value is `Hello, World!` |
| AC-3 | Stored greeting is `Custom greeting` | Visitor opens page | Large heading text is `Custom greeting` |
| AC-4 | Stored greeting is `Custom greeting` | Visitor opens page | Text input value is `Custom greeting` |

**Requirement GREETING-002 — Save non-empty greeting**

*As a* Visitor, *I want to* save changed greeting text, *so that* future page loads show my saved greeting.

Behaviour:

1. Visitor edits greeting text input.
2. Visitor submits form with Save button or equivalent form submit action.
3. If trimmed input is non-empty, saved greeting becomes trimmed input text.
4. Page updates large heading to saved greeting.
5. Text input value updates to saved greeting.
6. Polite message text reads `Saved.` after successful save.
7. Reloading page after successful save shows saved greeting in heading and input.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through story plan criteria that cite it.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Page shows stored greeting `Hello, World!` | Visitor enters `Hi pipeline` and saves | Large heading text becomes `Hi pipeline` |
| AC-2 | Page shows stored greeting `Hello, World!` | Visitor enters `Hi pipeline` and saves | Text input value becomes `Hi pipeline` |
| AC-3 | Save of `Hi pipeline` succeeds | Page displays result | Message text reads `Saved.` |
| AC-4 | Save of `Hi pipeline` succeeds | Visitor reloads page | Large heading text is `Hi pipeline` |
| AC-5 | Save of `Hi pipeline` succeeds | Visitor reloads page | Text input value is `Hi pipeline` |
| AC-6 | Page shows stored greeting `Hello, World!` | Visitor enters `  Hi pipeline  ` and saves | Saved greeting shown in heading is `Hi pipeline` |

**Requirement GREETING-003 — Reject empty greeting**

*As a* Visitor, *I want to* know when greeting text is empty, *so that* blank greeting is not saved by mistake.

Behaviour:

1. Visitor clears greeting input or enters only whitespace.
2. Visitor submits form.
3. Greeting is not changed.
4. Polite message text reads `Enter a greeting before saving.`.
5. Focus returns to greeting input.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through story plan criteria that cite it.

| # | Given | Stored greeting is `Hello, World!` | Then |
|---|---|---|---|
| AC-1 | Stored greeting is `Hello, World!` | Visitor clears input and saves | Message text reads `Enter a greeting before saving.` |
| AC-2 | Stored greeting is `Hello, World!` | Visitor clears input and saves | Large heading remains `Hello, World!` |
| AC-3 | Stored greeting is `Hello, World!` | Visitor enters spaces only and saves | Large heading remains `Hello, World!` |
| AC-4 | Stored greeting is `Hello, World!` | Visitor clears input and saves | Focus is on greeting input |

**Requirement GREETING-004 — Match approved screen**

*As a* Visitor, *I want to* use minimal approved greeting screen, *so that* page matches stakeholder-approved design.

Behaviour:

1. Page background is white `#FFFFFF`.
2. Text is black `#000000`.
3. Page content is centered in viewport.
4. Page contains one section only.
5. Section contains one large heading, one visually hidden label with text `Greeting`, one text input, one Save button, and one polite message area.
6. Save button has blue `#2563EB` background and border, white text, square corners, and label `Save`.
7. Input has white background, black text, black border, square corners, and autocomplete off.
8. Input and button show blue `#2563EB` focus outline when keyboard focused.
9. No navigation appears.
10. No animation or transition appears.
11. At viewport widths up to 520px, form controls stack vertically and Save button spans form width.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through story plan criteria that cite it.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Visitor opens page | Page renders | Background color is `#FFFFFF` and main text color is `#000000` |
| AC-2 | Visitor opens page | Page renders | Greeting section is centered in viewport |
| AC-3 | Visitor opens page | Page renders | Exactly one page section is present |
| AC-4 | Visitor opens page | Page renders | One large heading is present and labelled as page heading |
| AC-5 | Visitor opens page | Page renders | Text input has accessible label `Greeting` |
| AC-6 | Visitor opens page | Page renders | One Save button is present with text `Save` |
| AC-7 | Visitor opens page | Page renders | Save button background and border color are `#2563EB` |
| AC-8 | Visitor opens page | Page renders | Polite message area exists below form and is empty before submit |
| AC-9 | Visitor tabs to input or Save button | Element receives keyboard focus | Focus outline uses `#2563EB` |
| AC-10 | Visitor opens page | Page renders | No navigation landmark or navigation links are present |
| AC-11 | Visitor opens page | Page renders | No animation or transition is applied |
| AC-12 | Viewport width is 520px or less | Visitor opens page | Input and Save button are stacked vertically, and Save button spans form width |

**Failure, boundary and permission behaviour** — cases this function has from brief and approved design.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Greeting input is empty or whitespace only | Current greeting remains unchanged; message reads `Enter a greeting before saving.`; focus returns to input |
| Boundary | Greeting contains leading or trailing spaces | Spaces are trimmed before saving; displayed greeting uses trimmed value |
| Persistence | Visitor reloads after successful save | Last saved greeting appears in heading and input |
| Permission | Visitor is not signed in | Not applicable: no sign-in or roles exist; every visitor may view and save greeting |
| Error screen | Read or save cannot complete | Not applicable: approved design has no error state; API error envelope belongs in service contract |
| Loading screen | Greeting is being fetched or saved | Not applicable: approved design has no loading state |
| Empty screen | No greeting exists | Not applicable: stored greeting starts as `Hello, World!`; approved design has no empty state |
| Conflict | Two visitors save different greetings | Last successful save determines greeting shown on next load; no conflict UI exists in approved design |

**Data touched** — fields this function reads and writes, in product terms.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Initial value is `Hello, World!`; value saved by visitor must be non-empty after trimming; displayed exactly as saved after trimming |

## 5. Screens

Design is source of truth for appearance. Approved design shows one screen with one state. It includes: white page background, centered greeting section, large greeting heading, visually hidden `Greeting` label, text input prefilled with greeting, blue `Save` button, and reserved polite message area below form. The approved design also defines message text after client-side interactions: `Saved.` and `Enter a greeting before saving.`.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting page | `main > section.greeting-section` in approved `index.html` | GREETING-001, GREETING-002, GREETING-003, GREETING-004 | default |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Accessibility | Input has accessible label `Greeting`; message changes are exposed through polite live region; input and button are keyboard reachable; input and button have visible focus outline; text and button contrast ratios are at least 4.5:1. |
| Responsive | Page works from 320px viewport width upward with no horizontal page scroll; form stacks vertically at viewport widths up to 520px. |
| Persistence | A successfully saved greeting survives page reload and service restart because it is stored in PostgreSQL. |
| Localisation | UI copy is English: `Greeting`, `Save`, `Saved.`, `Enter a greeting before saving.`; no dates, numbers, or locale-specific formatting appear. |
| Privacy | Greeting text is public page content; no personal data, accounts, or secrets are collected by this module. |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL storage, for retaining one greeting across reloads and service restarts.
- **Depends on:** Go API, for reading and updating stored greeting.
- **Depends on:** Next.js frontend, for rendering approved page and handling visitor input.
- **Assumption:** One shared greeting exists for all visitors. If false, scope changes to accounts or per-session storage, which is outside approved plan.
- **Assumption:** Latest successful save wins when multiple visitors save. If false, scope changes to conflict handling UI, which approved design does not show.

| Open question | Proposed default | Who decides |
|---|---|---|
| — | None; no open stakeholder questions remain for this module. | — |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Persisted Editable Greeting | GREETING-001, GREETING-002, GREETING-003, GREETING-004 | `test-cases/persisted-editable-greeting.md` |
