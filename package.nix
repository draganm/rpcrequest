{ pkgs }:
pkgs.buildGoModule {
  pname = "rpcrequest";
  version = "0";
  src = ./.;
  vendorHash = "sha256-l8X2D+V9EWu1k6HF9/A63yt2+UtPc5WSE9NLsxawbtM=";
}
