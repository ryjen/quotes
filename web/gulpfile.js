import gulp from 'gulp';
import babel from 'gulp-babel';
import sass from 'gulp-dart-sass';
import postcss from 'gulp-postcss';
import concat from 'gulp-concat';
import imagemin from 'gulp-imagemin';
import uglify from 'gulp-uglify';
import changed from 'gulp-changed-in-place';
import cachebust from 'gulp-cache-bust';

import normalize from 'postcss-normalize';
import autoprefixer from 'autoprefixer';
import cssnano from 'cssnano';
import del from 'del';


const { src, dest, series, parallel, watch, lastRun } = gulp;

const paths = {
  js: {
    src: [
      'js/**.js'
    ],
    dest: './public/js',
  },
  css: {
    src: [
      'css/**/*.css',
      'sass/**/*.scss'
    ],
    dest: './public/css'
  },
  html: {
    src: [
      'public/**/*.html'
    ],
    dest: './public'
  },
  img: {
    src: [
      'img/**'
    ],
    dest: './public/img'
  },
  font: {
    src: [
      'font/**'
    ],
    dest: './public/font'
  },
  sourcemaps: "maps"
}

const js = () => src(paths.js.src, {
  since: lastRun(js),
  sourcemaps: true
})
  .pipe(changed({ firstPass: true }))
  .pipe(babel({
    presets: ['@babel/env']
  }))
  .pipe(concat('app.min.js'))
  .pipe(uglify())
  .pipe(dest(paths.js.dest, { sourcemaps: paths.sourcemaps }))

const css = () => src(paths.css.src, {
  since: lastRun(css),
  sourcemaps: true
})
  .pipe(changed({ firstPass: true }))
  .pipe(sass().on('error', sass.logError))
  .pipe(postcss([
    autoprefixer(),
    normalize(),
    cssnano(),
  ]))
  .pipe(concat('app.min.css'))
  .pipe(dest(paths.css.dest, { sourcemaps: paths.sourcemaps }))

const img = () => src(paths.img.src, {
  since: lastRun(img),
})
  .pipe(changed({ firstPass: true }))
  .pipe(imagemin({ silent: true }))
  .pipe(dest(paths.img.dest))

const font = () =>
  src(paths.font.src, {
    since: lastRun(font),
  })
    .pipe(changed({ firstPass: true }))
    .pipe(dest(paths.font.dest))

const html = () =>
  src(paths.html.src)
    .pipe(cachebust())
    .pipe(gulp.dest(paths.html.dest))

export const live = () =>
  watch('css', 'js', 'img', 'font')

export const clean = () =>
  del([
    './public/css/**',
    './public/js/**',
    './public/font/**',
    './public/img/**',
  ], { force: true })

export const build = parallel(js, css, img, font)

export default series(
  clean,
  build,
  html
)