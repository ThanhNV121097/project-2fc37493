import { PersistedEditableGreeting } from "@/components/PersistedEditableGreeting";
import { initialGreeting } from "@/lib/mock/persisted-editable-greeting";

export default function Page() {
  return (
    <main style={{ minHeight: "100vh", display: "grid", placeItems: "center", padding: "var(--space-24)" }}>
      <PersistedEditableGreeting initialGreeting={initialGreeting.text} />
    </main>
  );
}
