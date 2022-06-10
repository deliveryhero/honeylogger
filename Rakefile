task :command_exists, [:command] do |_, args|
  abort "#{args.command} doesn't exists" if `command -v #{args.command} > /dev/null 2>&1 && echo $?`.chomp.empty?
end

task :has_bumpversion do
    Rake::Task['command_exists'].invoke('bumpversion')
end
  
desc "default task"
task :default => [:run_server]

desc "run tests"
task :test do
  system "go test -p 1 -v -race ./..."
end

AVAILABLE_REVISIONS = %w[major minor patch].freeze
desc "bump version, default is: patch"
task :bump, [:revision] => [:has_bumpversion] do |_, args|
  args.with_defaults(revision: 'patch')
  unless AVAILABLE_REVISIONS.include?(args.revision)
    abort "Please provide valid revision: #{AVAILABLE_REVISIONS.join(',')}"
  end

  system "bumpversion #{args.revision}"
end

desc "run mockery"
task :mockery do
  system "mockery --dir ./logging --all --keeptree"
end