{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/release-23.05";

    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ ];
        };
      in
      with pkgs;
      {
        devShells.default = mkShell {
          buildInputs = [
            delve
            go
            gofumpt
            golangci-lint
            gopls
            gotest
            gotests
            gotestsum
            gotools
            jq
            openapi-generator-cli
          ];
        };

        formatter = pkgs.nixpkgs-fmt;
      }
    );
}
