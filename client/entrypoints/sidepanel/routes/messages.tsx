import { useEffect, useState } from 'react';
import { createFileRoute } from '@tanstack/react-router';

const API_BASE = 'http://localhost:8080';

type Message = { id: number; text: string; createdAt: string };

export const Route = createFileRoute('/messages')({
  component: MessagesPage,
});

function MessagesPage() {
  const [messages, setMessages] = useState<Message[]>([]);
  const [input, setInput] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  async function loadMessages() {
    setError(null);
    try {
      const res = await fetch(`${API_BASE}/api/messages`);
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      setMessages(await res.json());
    } catch (e) {
      setError(String(e));
    }
  }

  async function sendMessage() {
    if (!input.trim()) return;
    setLoading(true);
    setError(null);
    try {
      const res = await fetch(`${API_BASE}/api/messages`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text: input }),
      });
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const created: Message = await res.json();
      setMessages((prev) => [...prev, created]);
      setInput('');
    } catch (e) {
      setError(String(e));
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    loadMessages();
  }, []);

  return (
    <section>
      {error && <div className="error">⚠️ {error}</div>}
      <div className="row">
        <input
          value={input}
          placeholder="メッセージを入力"
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => e.key === 'Enter' && sendMessage()}
        />
        <button onClick={sendMessage} disabled={loading}>
          送信
        </button>
      </div>
      <button className="reload" onClick={loadMessages}>
        再取得
      </button>
      <ul className="messages">
        {messages.length === 0 && <li className="empty">まだメッセージがありません</li>}
        {messages.map((m) => (
          <li key={m.id}>
            <span className="text">{m.text}</span>
            <span className="time">{new Date(m.createdAt).toLocaleTimeString()}</span>
          </li>
        ))}
      </ul>
    </section>
  );
}
