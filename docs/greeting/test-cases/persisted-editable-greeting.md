# Test Cases — Persisted Editable Greeting

Module: `greeting`
Function: Persisted Editable Greeting
Story: `docs/greeting/stories/persisted-editable-greeting.md`
Service contract: `docs/architecture/services.md`

Risk level: Medium. This story is small, public, and has no roles or payments, but it writes shared persisted state and proves the full PostgreSQL → Go API → Next.js path.

## Page and interaction cases

**Scenario**: Initial greeting appears in heading
**Given**: No prior visitor-saved greeting exists and the stored singleton greeting is seeded as `Hello, World!`.
**When**: Visitor opens the greeting page.
**Then**: The large page heading text is exactly `Hello, World!`.
Traces: SC-1 (GREETING-001 AC-1)
Check: render_url

**Scenario**: Initial greeting appears in input
**Given**: No prior visitor-saved greeting exists and the stored singleton greeting is seeded as `Hello, World!`.
**When**: Visitor opens the greeting page.
**Then**: The greeting text input value is exactly `Hello, World!`.
Traces: SC-2 (GREETING-001 AC-2)
Check: render_url

**Scenario**: Custom stored greeting appears in heading
**Given**: Stored greeting is `Custom greeting`.
**When**: Visitor opens the greeting page.
**Then**: The large page heading text is exactly `Custom greeting`.
Traces: SC-3 (GREETING-001 AC-3)
Check: render_url

**Scenario**: Custom stored greeting appears in input
**Given**: Stored greeting is `Custom greeting`.
**When**: Visitor opens the greeting page.
**Then**: The greeting text input value is exactly `Custom greeting`.
Traces: SC-4 (GREETING-001 AC-4)
Check: render_url

**Scenario**: Saving non-empty greeting updates heading
**Given**: Page shows stored greeting `Hello, World!`.
**When**: Visitor replaces input text with `Hi pipeline` and clicks `Save`.
**Then**: The large page heading text becomes exactly `Hi pipeline`.
Traces: SC-5 (GREETING-002 AC-1)
Check: interact_page

**Scenario**: Saving non-empty greeting updates input
**Given**: Page shows stored greeting `Hello, World!`.
**When**: Visitor replaces input text with `Hi pipeline` and clicks `Save`.
**Then**: The greeting text input value becomes exactly `Hi pipeline`.
Traces: SC-6 (GREETING-002 AC-2)
Check: interact_page

**Scenario**: Successful save shows polite saved message
**Given**: Page shows stored greeting `Hello, World!`.
**When**: Visitor replaces input text with `Hi pipeline` and clicks `Save`.
**Then**: The polite message area text reads exactly `Saved.`.
Traces: SC-7 (GREETING-002 AC-3)
Check: interact_page

**Scenario**: Saved greeting persists in heading after reload
**Given**: Save of `Hi pipeline` succeeds.
**When**: Visitor reloads the page.
**Then**: The large page heading text is exactly `Hi pipeline`.
Traces: SC-8 (GREETING-002 AC-4)
Check: interact_page

**Scenario**: Saved greeting persists in input after reload
**Given**: Save of `Hi pipeline` succeeds.
**When**: Visitor reloads the page.
**Then**: The greeting text input value is exactly `Hi pipeline`.
Traces: SC-9 (GREETING-002 AC-5)
Check: interact_page

**Scenario**: Saved greeting is trimmed before display
**Given**: Page shows stored greeting `Hello, World!`.
**When**: Visitor replaces input text with `  Hi pipeline  ` and clicks `Save`.
**Then**: The large page heading text is exactly `Hi pipeline`, without leading or trailing spaces.
Traces: SC-10 (GREETING-002 AC-6)
Check: interact_page

**Scenario**: Empty input shows validation message
**Given**: Stored greeting is `Hello, World!` and page shows that greeting.
**When**: Visitor clears greeting input and clicks `Save`.
**Then**: The polite message area text reads exactly `Enter a greeting before saving.`.
Traces: SC-11 (GREETING-003 AC-1)
Check: interact_page

**Scenario**: Empty input does not change heading
**Given**: Stored greeting is `Hello, World!` and page shows that greeting.
**When**: Visitor clears greeting input and clicks `Save`.
**Then**: The large page heading text remains exactly `Hello, World!`.
Traces: SC-12 (GREETING-003 AC-2)
Check: interact_page

**Scenario**: Spaces-only input does not change heading
**Given**: Stored greeting is `Hello, World!` and page shows that greeting.
**When**: Visitor replaces input text with three spaces and clicks `Save`.
**Then**: The large page heading text remains exactly `Hello, World!`.
Traces: SC-13 (GREETING-003 AC-3)
Check: interact_page

**Scenario**: Empty input returns focus to greeting input
**Given**: Stored greeting is `Hello, World!` and page shows that greeting.
**When**: Visitor clears greeting input and clicks `Save`.
**Then**: Browser focus is on the greeting text input.
Traces: SC-14 (GREETING-003 AC-4)
Check: interact_page

## Approved screen cases

**Scenario**: Page uses white background and black text
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: Computed page background color is `#FFFFFF` and computed main text color is `#000000`.
Traces: SC-15 (GREETING-004 AC-1)
Check: measure_styles

**Scenario**: Greeting section is centered in viewport
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: `main > section.greeting-section` is centered in the viewport horizontally and vertically by computed layout.
Traces: SC-16 (GREETING-004 AC-2)
Check: measure_styles

**Scenario**: Page contains exactly one section
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: Exactly one `section` element is present on the page.
Traces: SC-17 (GREETING-004 AC-3)
Check: render_url

**Scenario**: Page has one large heading
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: Exactly one `h1` page heading is present, and its font size is larger than the input text font size.
Traces: SC-18 (GREETING-004 AC-4)
Check: measure_styles

**Scenario**: Greeting input has accessible label
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: The text input has accessible label text exactly `Greeting`.
Traces: SC-19 (GREETING-004 AC-5)
Check: render_url

**Scenario**: Save button appears once with exact text
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: Exactly one button is present and its visible text is exactly `Save`.
Traces: SC-20 (GREETING-004 AC-6)
Check: render_url

**Scenario**: Save button uses blue background and border
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: The `button` computed background color is `#2563EB` and computed border color is `#2563EB`.
Traces: SC-21 (GREETING-004 AC-7)
Check: measure_styles

**Scenario**: Polite message area is present and empty before submit
**Given**: Visitor opens the greeting page.
**When**: Page renders before any submit.
**Then**: A polite live message area exists below the form and its text content is empty.
Traces: SC-22 (GREETING-004 AC-8)
Check: render_url

**Scenario**: Keyboard focus outline uses blue on input and button
**Given**: Visitor opens the greeting page.
**When**: Visitor presses Tab until the greeting input receives focus, then presses Tab until the Save button receives focus.
**Then**: The focused input and focused button each show a computed focus outline using `#2563EB`.
Traces: SC-23 (GREETING-004 AC-9)
Check: interact_page

**Scenario**: Page has no navigation
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: No `nav` landmark is present and no navigation links are present.
Traces: SC-24 (GREETING-004 AC-10)
Check: render_url

**Scenario**: Page has no animation or transition
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: `body`, `main`, `main > section.greeting-section`, `h1`, `form`, `input`, `button`, and the polite message area have computed animation duration `0s` and transition duration `0s`.
Traces: SC-25 (GREETING-004 AC-11)
Check: measure_styles

**Scenario**: Small viewport stacks controls and full-width button
**Given**: Viewport width is 520px or less.
**When**: Visitor opens the greeting page.
**Then**: The greeting input and Save button are stacked vertically, and the Save button computed width equals the form content width.
Traces: SC-26 (GREETING-004 AC-12)
Check: measure_styles

## Service contract cases

**Scenario**: GET greeting returns current text JSON
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `GET /api/v1/greeting`.
**Then**: Response status is `200` and response body is exactly `{"text":"Hello, World!"}`.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting reports internal query failure with error envelope
**Given**: A greeting query fails while backend remains reachable.
**When**: Client requests `GET /api/v1/greeting`.
**Then**: Response status is `500` and response body is exactly `{"error":{"code":"INTERNAL","message":"Internal server error."}}`.
Traces: contract (GET /v1/greeting), error envelope
Check: manual

**Scenario**: PUT greeting stores trimmed text and returns current text JSON
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"text":"  Hi pipeline  "}`.
**Then**: Response status is `200` and response body is exactly `{"text":"Hi pipeline"}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects malformed JSON
**Given**: Backend is reachable.
**When**: Client sends `PUT /api/v1/greeting` with body `{` and `Content-Type: application/json`.
**Then**: Response status is `400` and response body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting), error envelope
Check: fetch_url

**Scenario**: PUT greeting rejects wrong JSON type
**Given**: Backend is reachable.
**When**: Client sends `PUT /api/v1/greeting` with JSON body `"Hi pipeline"`.
**Then**: Response status is `400` and response body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting), error envelope
Check: fetch_url

**Scenario**: PUT greeting rejects unknown request field
**Given**: Backend is reachable.
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"text":"Hi pipeline","extra":"ignored?"}`.
**Then**: Response status is `400` and response body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting), error envelope
Check: fetch_url

**Scenario**: PUT greeting rejects empty trimmed text
**Given**: Stored greeting is `Hello, World!`.
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"text":"   "}`.
**Then**: Response status is `422` and response body is exactly `{"error":{"code":"VALIDATION_FAILED","message":"Greeting must not be empty."}}`.
Traces: contract (PUT /v1/greeting), error envelope
Check: fetch_url

**Scenario**: PUT greeting reports internal query failure with error envelope
**Given**: A greeting update query fails while backend remains reachable.
**When**: Client sends `PUT /api/v1/greeting` with JSON body `{"text":"Hi pipeline"}`.
**Then**: Response status is `500` and response body is exactly `{"error":{"code":"INTERNAL","message":"Internal server error."}}`.
Traces: contract (PUT /v1/greeting), error envelope
Check: manual

**Scenario**: Last successful save wins
**Given**: Stored greeting is `Hello, World!`.
**When**: Client A successfully saves `First`, then Client B successfully saves `Second`, and Visitor opens the page.
**Then**: The large page heading text is exactly `Second` and the input value is exactly `Second`.
Traces: contract (PUT /v1/greeting), SRS conflict behavior
Check: interact_page
