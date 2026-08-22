{
  description = "WXT + React sidepanel extension dev environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.nodejs_22
            pkgs.pnpm
          ];

          shellHook = ''
            echo "🧩 sidepanel-extension dev shell"
            echo "  node: $(node --version) / pnpm: $(pnpm --version)"
            echo "  → pnpm install && pnpm dev"
          '';
        };
      });
}
