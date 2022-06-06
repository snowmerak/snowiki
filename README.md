# snowiki

snowiki는 해시태그 기반의 간단한 정적 웹 위키입니다. 

## 설치

`go install github.com/snowmerak/snowiki@latest`로 설치할 수 있습니다.

## 사용법

프로젝트 폴더를 하나 생성하고, 내부에 `src` 폴더를 만듭니다.

`src` 폴더 내에 마크다운 문서(`.md`)로 문서를 작성합니다.

문서 작성 중 해시태그를 작성(`#hashtag`)하시면 자동으로 문서가 해시태그에 연결됩니다.

문서를 작성하시고, 프로젝트 폴더에서 `snowiki`를 실행하시면 `public` 폴더에 `html` 파일이 생성됩니다.
