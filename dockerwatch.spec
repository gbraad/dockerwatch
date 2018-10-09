%global executable_name dockerwatch

# Generate devel rpm
%global with_devel 1
# Build project from bundled dependencies
%global with_bundled 0
# Build with debug info rpm
%global with_debug 1
# Run tests in check section
%global with_check 1
# Generate unit-test rpm
%global with_unit_test 1

%if 0%{?with_debug}
%global _dwz_low_mem_die_limit 0
%else
%global debug_package   %{nil}
%endif

%if ! 0%{?gobuild:1}
%define gobuild(o:) go build -ldflags "${LDFLAGS:-} -B 0x$(head -c20 /dev/urandom|od -An -tx1|tr -d ' \\n')" -a -v -x %{?**}; 
%endif

%global provider        github
%global provider_tld    com
%global project         gbraad
%global repo            dockerwatch
# https://github.com/gbraad/dockerwatch
%global provider_prefix %{provider}.%{provider_tld}/%{project}/%{repo}
%global import_path     %{provider_prefix}

Name:           %{repo}
Version:        0.0.1
Release:        2%{?dist}
Summary:        dockerwatch: Execute commands on new containers
License:        ASL 2.0
URL:            https://%{provider_prefix}
Source0:        https://%{provider_prefix}/archive/%{version}.tar.gz

# If go_compiler is not set to 1, there is no virtual provide. Use golang instead.
BuildRequires:  %{?go_compiler:compiler(go-compiler)}%{!?go_compiler:golang}

# handle license on el{6,7}: global must be defined after the License field above
%{!?_licensedir: %global license %doc}

%description
%{summary}

%prep
%setup -q -n %{name}-%{version}

%build
mkdir -p src/%{import_path}
rmdir src/%{import_path}
ln -s ../../../ src/%{import_path}

%if ! 0%{?with_bundled}
export GOPATH=$(pwd):%{gopath}
%else
echo "Unable to build from bundled deps. No Godeps nor vendor directory"
exit 1
%endif

%gobuild -o bin/%{executable_name} %{import_path}/cmd/%{executable_name}

%install
install -d -p %{buildroot}%{_bindir}
install -m 755 bin/%{executable_name} %{buildroot}%{_bindir}

%check
%if 0%{?with_check}
! %{buildroot}%{_bindir}/%{executable_name} --help
%endif

%files
%doc README.md
%license LICENSE
%{_bindir}/%{executable_name}

%changelog
* Wed Oct 10 2018 Gerard Braad <me@gbraad.nl> 0.0.1
- Initial version
