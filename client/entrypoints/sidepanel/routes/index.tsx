import { useEffect, useState } from 'react';
import { createFileRoute } from '@tanstack/react-router';

const API_BASE = 'http://localhost:8080';

type Health = { status: string; time: string };

export const Route = createFileRoute('/')({
  component: HomePage,
});

function HomePage() {
  const [health, setHealth] = useState<Health | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function loadHealth() {
    setError(null);
    try {
      const res = await fetch(`${API_BASE}/api/health`);
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      setHealth(await res.json());
    } catch (e) {
      setError(String(e));
    }
  }

  useEffect(() => {
    loadHealth();
  }, []);

  return (
    <section>
      <div className={`badge ${health ? 'ok' : 'ng'}`}>
        {health ? `API: ${health.status}` : 'API: 未接続'}
      </div>
      {health && <p className="time">サーバ時刻: {health.time}</p>}
      {error && <div className="error">⚠️ {error}</div>}
      <button className="reload" onClick={loadHealth}>
        再チェック
      </button>
    </section>
  );
}
