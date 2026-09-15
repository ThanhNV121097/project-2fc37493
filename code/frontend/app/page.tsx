import { PersistedEditableGreeting } from "@/components/PersistedEditableGreeting";
import { readGreeting } from "@/lib/persisted-editable-greeting";

export default async function Page() {
  const greeting = await readGreeting(process.env.API_ORIGIN ?? "http://backend:8080");

  return <PersistedEditableGreeting initialGreeting={greeting.text} />;
}
