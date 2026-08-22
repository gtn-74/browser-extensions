import { createRootRoute, Link, Outlet } from '@tanstack/react-router';

export const Route = createRootRoute({
  component: RootLayout,
});

function RootLayout() {
  return (
    <div className="app">
      <header>
        <h1>Sidepanel Lab</h1>
        <nav className="nav">
          <Link to="/" activeProps={{ className: 'active' }}>
            Home
          </Link>
          <Link to="/messages" activeProps={{ className: 'active' }}>
            Messages
          </Link>
        </nav>
      </header>
      <Outlet />
    </div>
  );
}
