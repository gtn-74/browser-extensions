{
  description = "Go sidepanel API dev environment";

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
            pkgs.go
            pkgs.gopls
            pkgs.gotools
          ];

          shellHook = ''
            # mise などが設定する GOROOT を消し、nix の go に自分の GOROOT を使わせる
            unset GOROOT
            echo "🐹 sidepanel-api dev shell"
            echo "  go: $(go version)"
            echo "  → go run ."
          '';
        };
      });
}
