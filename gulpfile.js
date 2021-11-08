const gulp = require('gulp')
const babel = require('gulp-babel')
const sourcemaps = require('gulp-sourcemaps')
// const nodemon = require('gulp-nodemon')
const dotenv = require('dotenv')
const Cache = require('gulp-file-cache')
const run = require('gulp-run-command').default

dotenv.config()

const path = {
  env: './.env',
  test: "./test/**/*.js",
  js: './app/**/*.js',
  storage: './storages/cache/',
  apiController: './app/Controllers/api/**/*.js'
}
const cache = new Cache()
cache.clear()

//goi hàm chuyển es6 -> es5
gulp.task('es6', () => {
  return gulp.src([path.js]).pipe(sourcemaps.init()).pipe(cache.filter()).pipe(babel({
    presets: ['@babel/preset-env'],
    plugins: [
      [
        '@babel/plugin-transform-runtime',
        {'corejs': false, 'helpers': true, 'regenerator': true, 'useESModules': false},
      ],
      ['@babel/plugin-proposal-class-properties'],
      ['@babel/plugin-proposal-private-methods'],
      ['@babel/plugin-proposal-object-rest-spread', {'loose': true, 'useBuiltIns': true}],
    ],
  })).pipe(cache.cache()).pipe(gulp.dest(path.storage + "app"))
})

gulp.task('test', () => {
  return gulp.src([path.test]).pipe(sourcemaps.init()).pipe(cache.filter()).pipe(babel({
    presets: ['@babel/preset-env'],
    plugins: [
      [
        '@babel/plugin-transform-runtime',
        {'corejs': false, 'helpers': true, 'regenerator': true, 'useESModules': false},
      ],
      ['@babel/plugin-proposal-class-properties'],
      ['@babel/plugin-proposal-private-methods'],
      ['@babel/plugin-proposal-object-rest-spread', {'loose': true, 'useBuiltIns': true}],
    ],
  })).pipe(cache.cache()).pipe(gulp.dest(path.storage + "test"))
})
//build documents
gulp.task('document', async () => run('npm run document')())

//xem thay doi js để build lại cache es5
gulp.task('watchFiles', () => {
  gulp.watch([path.js], gulp.series('es6'))
  gulp.watch([path.test], gulp.series('test'))
  gulp.watch([path.apiController], gulp.series('document'))
})

let runStack
if (process.env.NODE_ENV === 'development') {
  runStack = gulp.series( gulp.parallel('es6', 'test'), gulp.parallel('watchFiles'))
} else {
  runStack = gulp.series( gulp.parallel('es6', 'test'))
}

gulp.task('build', gulp.series('es6','test'))
gulp.task('default', runStack)

