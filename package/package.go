package ctff

import (
	cbase "github.com/jurgen-kluft/cbase/package"
	"github.com/jurgen-kluft/ccode/denv"
	ccore "github.com/jurgen-kluft/ccore/package"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

const (
	repo_path = "github.com\\jurgen-kluft\\"
	repo_name = "ctff"
)

func GetPackage() *denv.Package {
	name := repo_name

	// Dependencies
	basepkg := cbase.GetPackage()
	corepkg := ccore.GetPackage()
	unittestpkg := cunittest.GetPackage()

	// The main (ctff) package
	mainpkg := denv.NewPackage(name)
	mainpkg.AddPackage(basepkg)
	mainpkg.AddPackage(corepkg)

	// library
	mainlib := denv.SetupCppLibProject(name, repo_path+name)
	mainlib.AddDependencies(basepkg.GetMainLib()...)
	mainlib.AddDependencies(corepkg.GetMainLib()...)

	// unittest project
	maintest := denv.SetupCppTestProject(name+"_test", repo_path+name)
	maintest.AddDependencies(unittestpkg.GetMainLib()...)
	maintest.AddDependency(mainlib)

	// the cli application
	mainapp := denv.SetupDefaultCppCliProject(name+"_app", repo_path+name)
	mainapp.Dependencies = append(mainapp.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
