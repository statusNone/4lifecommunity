#!/usr/bin/env bash
set -euo pipefail

OUT="$(cd "$(dirname "$0")/../static/img" && pwd)"
mkdir -p "$OUT"

BASE="https://images.squarespace-cdn.com/content/v1/65bc113d1f64d44d4d635cbd"

# URL-path | local filename
IMAGES=(
  "4f8b612a-bf8c-4c35-86ef-8f0d9fcfc766/port_townsend_grant_office.png|hero-bg.png"
  "b552975f-8acf-4399-a236-d505ce10880a/April_Cutout_homepage.png|april-cutout-home.png"
  "6cfcbbd4-fa28-4241-9811-c0db00b8e376/colombia_misty.png|colombia-misty.png"
  "497b0342-120e-4881-85a7-0daf3bc0be81/%239+-+April+Bartlett+-+HeadshotPro.png|april-headshot.png"
  "1754957268111-RPPW41EHBCVQB3DW8T7J/image-asset.jpeg|adaptogenic-room.jpeg"
  "5efdfd24-7596-40e9-a31b-28fc7f700058/Bio%29_nomad.jpg|bioregional-nomad.jpg"
  "b96cbeb3-e7db-40da-ba70-026a99a27e0a/colombia.png|colombia.png"
  "9e84f10c-d7fd-4d62-9a88-66daa79c514a/4life_footer_tint_v2.png|footer-tint.png"
  "cf901dac-3524-43a2-b06f-61a9bdce7403/4life_logo_light.png|logo-light.png"
  "c49f18ec-baae-4feb-9e9e-81e785a01011/4life_logo_medium.png|logo.png"
  "447ad3b3-b435-4741-8c2c-5b9554fadcd7/4life+April+Bartlet+2025+-+Artboard+13.jpg|og-image.jpg"
  "4afefb4d-650e-4f39-badb-61eabc315f3e/Colombia_blur.png|blog-hero-bg.png"
  "da9e5b0f-895a-4027-9ed0-21afa721e39e/April_Cutout_Casual_V3.png|april-cutout-casual.png"
  "1757986063818-V5SSGKY621KZAT343CU9/unsplash-image-DNkoNXQti3c.jpg|post-prosocial.jpg"
  "1755040475811-QF6V920A2AUJEWRB43RK/Ashley_Skeeter.jpg|post-soil.jpg"
  "1755041363319-519UOQBG6SZ0RCZRN3EH/unsplash-image-CZJp1S4bZos.jpg|post-garn.jpg"
  "1754936292173-SLS8YOTAPSMZD0NI9XT2/Colombia_yoga.png|post-education.png"
  "1754936499357-GNBZ7XKLNILM1907XHU0/Colombia_Regenerative_Storytelling.png|post-storytelling.png"
  "1754936803784-7KIRT16NHRRYIPB2HBGJ/unsplash-image-kKwa4A1p__Q.jpg|post-backpacking.jpg"
  "1754936927648-6F1NUNYMZFG5BDMQIFE8/image-asset.jpeg|post-poetry.jpeg"
  "1761766756501-9622U4VFGMBCS6I54FZ2/gen_catatumbo_river_colombia.png|post-catatumbo.png"
  "c5fda57b-db8d-419b-ac58-f376b21a08e2/april_portrait.png|april-portrait.png"
  "21865936-a53f-43ef-b6cd-d527bd68ad99/Chimacum+Creek+Farm.jpg|chimacum-farm.jpg"
  "648c2955-1dc6-44d2-bc3c-2b6640737ed3/FinnriverFarmView_3.jpg|finnriver.jpg"
  "8ccfed75-5d51-44ac-8f0e-7194d1249143/Food+Bank+Gardens.jpg|food-bank-gardens.jpg"
  "da654444-5bbb-4072-8e7c-51d56b3b4197/Small+Habitat+Restoration.jpg|habitat-restoration.jpg"
  "f329e2b6-f3da-4bd3-86d6-a879e7e787be/April_Bartlett_Loft.jpg|april-loft.jpg"
  "440e05e0-c1cb-4fbb-8f62-c059150553cc/AstroYoga_Masthead.png|astroyoga-masthead.png"
  "0770b09a-e2e3-4133-b0f0-9108585184b8/Icon_Lotus.png|icon-lotus.png"
  "58bb241a-ba33-4904-82c6-63166af87408/Icon_yoga.png|icon-yoga.png"
  "db97edfe-7974-44fc-ae82-6af412fdf7f9/Icon_BodySpirit.png|icon-bodyspirit.png"
  "1771289378243-B18ZX9KZ4ID7BGR4BXZ2/unsplash-image-EwKXn5CapA4.jpg|adapt-hero.jpg"
  "1771289419697-WKW1SNKXCLHUIVYW5TKV/unsplash-image-mNGaaLeWEp0.jpg|adapt-body.jpg"
  "1759203046531-AZTOFU9HSG43RCV7GUE6/unsplash-image-qv5qEo_R0fw.jpg|catatumbo-hero.jpg"
  "f30451fc-3744-489f-b5b0-418745c01257/WhatsApp+Image+2025-09-26+at+3.11.34+PM+%281%29.jpeg|catatumbo-river.jpeg"
  "3711d864-b0c4-4555-9ab6-09021b22588d/Ecocide+on+the+Catatumbo+River+Oil+Silences+the+River+but+with+Your+Help+We+Will+Be+Heard+%282%29.png|catatumbo-ecocide.png"
  "87d99990-c02c-41b5-b2d9-b00942df0af1/492229496_122118548582825661_6091050871111386469_n.jpg|catatumbo-fisher1.jpg"
  "c04930c6-e5dc-4d58-ae46-ab541be4487f/546187412_1346459653909610_7624215926333231224_n.jpg|catatumbo-fisher2.jpg"
  "f29e4615-a754-431c-8fff-fec7d6fd8488/Bioregional_nomad.jpg|bioregional-nomads-hero.jpg"
  "4b93d81d-8206-47a9-9b78-f75a4a2e567b/WhatsApp+Image+2025-06-28+at+6.44.05+PM.jpeg|susurros-children.jpeg"
  "1754946327569-C0B2QH72LE4XCEDJ31YE/unsplash-image-T0O6qGzu03I.jpg|susurros-hillside.jpg"
  "9702af77-73ea-48de-a6cc-f2443a8e80c1/WhatsApp+Image+2025-06-28+at+7.49.02+PM.jpeg|susurros-scenes.jpeg"
)

for entry in "${IMAGES[@]}"; do
  src="${entry%%|*}"
  dst="${entry##*|}"
  if [ -f "$OUT/$dst" ]; then
    echo "skip $dst"
    continue
  fi
  echo "fetch $dst"
  curl -sL --max-time 60 "$BASE/$src" -o "$OUT/$dst"
done

echo "done"
