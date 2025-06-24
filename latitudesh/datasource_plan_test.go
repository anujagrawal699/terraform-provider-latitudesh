package latitudesh

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testPlanName = "c2.small.x86"

func TestAccPlan_Basic(t *testing.T) {

	recorder, teardown := createTestRecorder(t)
	defer teardown()
	testAccProviders["latitudesh"].ConfigureContextFunc = testProviderConfigure(recorder)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccTokenCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPlanBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.latitudesh_plan.test", "slug", testPlanName),
				),
			},
		},
	})
}

func testAccCheckPlanBasic() string {
	return fmt.Sprintf(`
data "latitudesh_plan" "test" {
	slug = "%s"
}
`,
		testPlanName,
	)
}
