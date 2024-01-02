{
  description = "A Nix-flake-based Go 1.17 development environment";

 inputs = {
    templ.url = "github:a-h/templ";
    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1.*.tar.gz";
  };

  outputs = { self, nixpkgs, templ }@inputs:
    let
      goVersion = 20; # Change this to update the whole stack
      overlays = [ (final: prev: { go = prev."go_1_${toString goVersion}"; }) ];
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = import nixpkgs { inherit overlays system; };
      });
        templ = system: inputs.templ.packages.${system}.templ;
    in
    {
      devShells = forEachSupportedSystem ({ pkgs }: {
        default = pkgs.mkShell {
          packages = with pkgs; [

            # go stuff
            go
            gotools # goimports, godoc, etc.
            golangci-lint # https://github.com/golangci/golangci-lint


            # Dependencies
            gnumake
            sqlite
            sqlc
            (templ system)
          ];
        };
      });
    };
}
