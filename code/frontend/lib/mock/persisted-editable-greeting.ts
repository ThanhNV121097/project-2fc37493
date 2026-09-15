export type Greeting = {
  text: string;
};

const storageKey = "persisted-editable-greeting";

export const initialGreeting: Greeting = {
  text: "Hello, World!",
};

export function readMockGreeting(): Greeting {
  if (typeof window === "undefined") {
    return initialGreeting;
  }

  return { text: window.localStorage.getItem(storageKey) ?? initialGreeting.text };
}

export function saveMockGreeting(text: string): Greeting {
  const greeting = { text };
  window.localStorage.setItem(storageKey, greeting.text);
  return greeting;
}
