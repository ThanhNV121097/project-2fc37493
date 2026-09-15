export type Greeting = {
  text: string;
};

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "/api";

export async function readGreeting(origin = apiBase): Promise<Greeting> {
  const response = await fetch(`${origin}/v1/greeting`, { cache: "no-store" });

  if (!response.ok) {
    throw new Error("Failed to load greeting.");
  }

  return response.json();
}

export async function saveGreeting(text: string): Promise<Greeting> {
  const response = await fetch(`${apiBase}/v1/greeting`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ text }),
  });

  if (!response.ok) {
    throw new Error("Failed to save greeting.");
  }

  return response.json();
}
