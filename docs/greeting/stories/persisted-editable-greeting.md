# Story — Persisted Editable Greeting

Module: `greeting`
Plan item: Persisted Editable Greeting
SRS: `docs/greeting/SRS.md`
Design system: `design/design-system.md`
Architecture: `docs/architecture/overview.md`

## User story

As a Visitor, I want to view and update the shared greeting, so that the page reflects the latest greeting stored in PostgreSQL after reload.

## In scope

- Seed PostgreSQL with one shared greeting row whose initial text is `Hello, World!`.
- Serve current greeting through Go API.
- Accept valid greeting updates through Go API and persist trimmed text in PostgreSQL.
- Render one Next.js page section showing current greeting as large heading.
- Render one text input prefilled with current greeting and one `Save` button.
- Let any Visitor save non-empty greeting text and see heading, input, and polite message update.
- Preserve saved greeting after page reload and backend restart.
- Apply approved minimal visual design and accessibility behavior.

## Out of scope

- Sign-in, roles, or permissions; every Visitor may view and save.
- Navigation, extra sections, extra pages, or marketing content.
- Multiple greetings, per-visitor greetings, greeting history, audit logs, or conflict UI.
- External services, notifications, analytics, or email.
- Loading, empty, or API error screens; approved design has no such states.
- Greeting maximum length limit; architecture notes say do not invent one.
- Hover, active, or disabled visual states beyond native browser behavior; approved design documents default and focus states only.

## UI scope

This story owns the full approved greeting page: `main > section.greeting-section`. Screen contains exactly one centered section with one large `h1`, visually hidden `Greeting` label, one text input, one `Save` submit button, and one polite message area below the form. Default message area is empty. Successful save message reads `Saved.`. Empty-input validation message reads `Enter a greeting before saving.`. Page uses white background `#FFFFFF`, black text `#000000`, blue `#2563EB` button and focus outline, square corners, no animation, no navigation, and compact layout stacking controls at widths up to 520px.

## Acceptance criteria

### Stored greeting display

- SC-1 [GREETING-001 AC-1]: Given no prior visitor-saved value exists, when Visitor opens the page, large heading text is `Hello, World!`.
- SC-2 [GREETING-001 AC-2]: Given no prior visitor-saved value exists, when Visitor opens the page, text input value is `Hello, World!`.
- SC-3 [GREETING-001 AC-3]: Given stored greeting is `Custom greeting`, when Visitor opens the page, large heading text is `Custom greeting`.
- SC-4 [GREETING-001 AC-4]: Given stored greeting is `Custom greeting`, when Visitor opens the page, text input value is `Custom greeting`.

### Saving greeting

- SC-5 [GREETING-002 AC-1]: Given page shows stored greeting `Hello, World!`, when Visitor enters `Hi pipeline` and saves, large heading text becomes `Hi pipeline`.
- SC-6 [GREETING-002 AC-2]: Given page shows stored greeting `Hello, World!`, when Visitor enters `Hi pipeline` and saves, text input value becomes `Hi pipeline`.
- SC-7 [GREETING-002 AC-3]: Given save of `Hi pipeline` succeeds, when page displays result, message text reads `Saved.`.
- SC-8 [GREETING-002 AC-4]: Given save of `Hi pipeline` succeeds, when Visitor reloads page, large heading text is `Hi pipeline`.
- SC-9 [GREETING-002 AC-5]: Given save of `Hi pipeline` succeeds, when Visitor reloads page, text input value is `Hi pipeline`.
- SC-10 [GREETING-002 AC-6]: Given page shows stored greeting `Hello, World!`, when Visitor enters `  Hi pipeline  ` and saves, saved greeting shown in heading is `Hi pipeline`.

### Empty greeting rejection

- SC-11 [GREETING-003 AC-1]: Given stored greeting is `Hello, World!`, when Visitor clears input and saves, message text reads `Enter a greeting before saving.`.
- SC-12 [GREETING-003 AC-2]: Given stored greeting is `Hello, World!`, when Visitor clears input and saves, large heading remains `Hello, World!`.
- SC-13 [GREETING-003 AC-3]: Given stored greeting is `Hello, World!`, when Visitor enters spaces only and saves, large heading remains `Hello, World!`.
- SC-14 [GREETING-003 AC-4]: Given stored greeting is `Hello, World!`, when Visitor clears input and saves, focus is on greeting input.

### Approved screen

- SC-15 [GREETING-004 AC-1]: When Visitor opens the page, background color is `#FFFFFF` and main text color is `#000000`.
- SC-16 [GREETING-004 AC-2]: When Visitor opens the page, greeting section is centered in the viewport.
- SC-17 [GREETING-004 AC-3]: When Visitor opens the page, exactly one page section is present.
- SC-18 [GREETING-004 AC-4]: When Visitor opens the page, one large heading is present and labelled as page heading.
- SC-19 [GREETING-004 AC-5]: When Visitor opens the page, text input has accessible label `Greeting`.
- SC-20 [GREETING-004 AC-6]: When Visitor opens the page, one Save button is present with text `Save`.
- SC-21 [GREETING-004 AC-7]: When Visitor opens the page, Save button background and border color are `#2563EB`.
- SC-22 [GREETING-004 AC-8]: When Visitor opens the page, polite message area exists below form and is empty before submit.
- SC-23 [GREETING-004 AC-9]: When Visitor tabs to input or Save button, focused element has focus outline using `#2563EB`.
- SC-24 [GREETING-004 AC-10]: When Visitor opens the page, no navigation landmark or navigation links are present.
- SC-25 [GREETING-004 AC-11]: When Visitor opens the page, no animation or transition is applied.
- SC-26 [GREETING-004 AC-12]: Given viewport width is 520px or less, when Visitor opens the page, input and Save button stack vertically and Save button spans form width.

## Dependencies

- PostgreSQL is available and reachable through `DATABASE_URL`.
- Backend service runs migrations at startup and seeds shared greeting row id `1` idempotently.
- Go API uses `/v1/...` routes and validates trimmed non-empty greeting text at API boundary.
- Next.js frontend can reach backend through configured API origin.
- No external accounts, credentials, or stakeholder decisions required.
