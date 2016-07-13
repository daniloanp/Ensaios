package test
//
//import (
//	"testing"
//	"github.com/daniloanp/Ensaios/application/backend/model"
//	"github.com/daniloanp/Ensaios/application/backend/app"
//	"database/sql"
//	"fmt"
//)
//
//type testModule struct{
//	*testing.T
//	model.Module
//}
//
//const (
//	logPadding = "\t => "
//)
//
//type testModuleCase struct {
//	in              *model.ModuleData
//	parentModuleInx sql.NullInt64
//	expectedErr     bool
//	err             error
//	testFunc        func(in *model.ModuleData, out *model.ModuleData, err error) bool
//}
//
//func makeValidateFunc(in *model.ModuleData, out *model.ModuleData, err error) bool {
//	return false
//}
////TestCase simple creating test
//func (t *testModule) TestCreate() {
//	t.Log("Testing Module.Create!")
//	cases := []*testModuleCase{//no parent
//		{
//			expectedErr: true,
//			err: model.ErrDataIsNil,
//		},
//		{
//			in: &model.ModuleData{Name:""}, // faker
//		},
//		{
//			in: &model.ModuleData{Name:"name"}, // faker
//			parentModuleInx: sql.NullInt64{1, true},
//		},
//		{
//			in: &model.ModuleData{Name:"should not pass"},
//			expectedErr: true,
//			parentModuleInx: sql.NullInt64{1, true},
//		},
//		{
//			in: &model.ModuleData{Name:"should_not_pass"},
//			expectedErr: true,
//			parentModuleInx: sql.NullInt64{1, true},
//		},
//		{
//			in: &model.ModuleData{Name:"ShouldNotPass"},
//			expectedErr: true,
//			parentModuleInx: sql.NullInt64{1, true},
//		},
//		{
//			in: &model.ModuleData{Name:"should-pass"}, //
//			parentModuleInx: sql.NullInt64{1, true},
//		},
//		{
//			in: &model.ModuleData{Name:"-should-not-pass"}, //
//			expectedErr: true,
//		},
//		{
//			in: &model.ModuleData{Name:"module-level"}, //
//			parentModuleInx: sql.NullInt64{1, true},
//		},
//		{
//			in: &model.ModuleData{Name:"fk-error", ParentModuleID:sql.NullInt64{-1, true}}, //Bad case
//			expectedErr: true,
//		},
//	}
//
//	for inx, cs := range cases {
//		t.Logf("==> Case %d:", inx)
//		if cs.parentModuleInx.Valid {
//			if int64(inx) + 1 <= cs.parentModuleInx.Int64 || cs.parentModuleInx.Int64 < 0 {
//				t.Fatal("Can't recover: you wrote your test bad")
//			} else {
//				parent := cases[int(cs.parentModuleInx.Int64)]
//				if parent.in == nil {
//					t.Fatal("Can't recover: you wrote your test bad")
//				} else {
//					cs.in.ParentModuleID = sql.NullInt64{Int64:parent.in.ID, Valid: true}
//				}
//			}
//		}
//		var in model.ModuleData
//		if cs.in != nil {
//			in = *cs.in
//		}
//		err := t.Create(cs.in)
//		out := cs.in
//		if err != nil {
//			if cs.expectedErr && (cs.err == nil || cs.err == err) {
//				t.Logf(logPadding+  "Expected error: %q", err)
//			} else {
//				t.Errorf(logPadding + "Failed: Error: %q ", err)
//			}
//		} else if cs.expectedErr {
//			t.Errorf(logPadding + "Failed: Some error was expected!")
//		} else {
//			// since we pass a reference, the function can change the content of model
//			if out.ID == in.ID {
//				//A call to create always insert a new line
//				t.Errorf(logPadding + "Same id (int64:%d) returned, expecting a different one", in.ID)
//				t.Compare(&in, out)
//			}
//
//			//now, get from db and test retrieved objects
//			retrivied, err := t.GetByID(out.ID)
//			if err != nil && retrivied != nil {
//
//			}
//			fmt.Printf("%v\n", retrivied)
//			if err != nil {
//				t.Errorf(logPadding + "Get failed when created are supposed to worked!")
//			} else {
//				t.Compare(out, retrivied)
//				if out.ID != retrivied.ID {
//					t.Errorf(logPadding + "Ids differing from original!")
//				}
//			}
//		}
//	}
//}
//
//func (t *testModule) Compare(in, out *model.ModuleData) {
//	if in.Name != out.Name {
//		t.Errorf(logPadding +  "The resulted name is diferent from the expected one. '%q' = '%q'", in.Name, out.Name)
//	}
//	if in.ParentModuleID != out.ParentModuleID {
//		t.Errorf(logPadding +  "The resulted parent_id is diferent from the expected one. %d = %d ", in.ParentModuleID, out.ParentModuleID)
//	}
//	if out.ID <= 0 {
//		t.Errorf(logPadding + "An invalid ID was returned. id == %d", out.ID)
//	}
//}
//
//func (t *testModule) TestDeleteByID() {
//	t.Log("Testing Module.DeleteByID!")
//}
//
//func (t *testModule) TestUpdate() {
//	t.Log("Testing Module.Update!")
//}
//
//func (t *testModule) TestSequence() {
//	t.Log("Testing Module.Update!")
//}
//
//func TestModule(t *testing.T) {
//	tester := &testModule{t, app.Db().Module()}
//	tester.TestCreate()
//	if !t.Failed() {
//		tester.TestDeleteByID()
//		tester.TestUpdate()
//	}
//}
//
//
