import { PersistedEditableGreeting } from "@/components/PersistedEditableGreeting";
import { initialGreeting } from "@/lib/mock/persisted-editable-greeting";

export default function Page() {
  return <PersistedEditableGreeting initialGreeting={initialGreeting.text} />;
}
