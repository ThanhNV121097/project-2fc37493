"use client";

import { FormEvent, useEffect, useRef, useState } from "react";
import { readMockGreeting, saveMockGreeting } from "@/lib/mock/persisted-editable-greeting";
import styles from "./PersistedEditableGreeting.module.css";

export type PersistedEditableGreetingProps = {
  initialGreeting: string;
};

export function PersistedEditableGreeting({ initialGreeting }: PersistedEditableGreetingProps) {
  const [greeting, setGreeting] = useState(initialGreeting);
  const [inputValue, setInputValue] = useState(initialGreeting);
  const [message, setMessage] = useState("");
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    const storedGreeting = readMockGreeting().text;
    setGreeting(storedGreeting);
    setInputValue(storedGreeting);
  }, []);

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const nextGreeting = inputValue.trim();

    if (!nextGreeting) {
      setMessage("Enter a greeting before saving.");
      inputRef.current?.focus();
      return;
    }

    const savedGreeting = saveMockGreeting(nextGreeting);
    setGreeting(savedGreeting.text);
    setInputValue(savedGreeting.text);
    setMessage("Saved.");
  }

  return (
    <main className={styles.main}>
      <section className={styles.section} aria-labelledby="greeting-heading">
        <h1 id="greeting-heading" className={styles.heading}>{greeting}</h1>
        <form className={styles.form} onSubmit={handleSubmit} noValidate>
          <label className="visually-hidden" htmlFor="greeting-input">Greeting</label>
          <input
            ref={inputRef}
            id="greeting-input"
            name="greeting"
            type="text"
            value={inputValue}
            autoComplete="off"
            required
            onChange={(event) => setInputValue(event.target.value)}
            className={styles.input}
          />
          <button type="submit" className={styles.button}>Save</button>
        </form>
        <p className={styles.message} aria-live="polite">{message}</p>
      </section>
    </main>
  );
}
