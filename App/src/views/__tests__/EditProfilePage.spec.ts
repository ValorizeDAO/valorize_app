import { shallowMount, DOMWrapper } from "@vue/test-utils"
import { mocked } from 'ts-jest/utils'
import EditProfilePage from "@/views/EditProfilePage.vue"
import Vuex, { Store } from 'vuex'
import auth from '../../services/authentication'

jest.mock('../../services/authentication', () => ({}))
const mockedAuth = mocked(auth, true)

describe("Dashboard.vue", () => {
  let store: Store<any>
  beforeEach(() => {
    //new vuex store with getters and mocks
    store = new Vuex.Store({
      getters: {
        profileImage: jest.fn(() => 'fake_url.jpg')
      }
    })
  })
  it("mounts", () => {
    const wrapper = shallowMount(EditProfilePage, {
      global: {
        plugins: [store],
      },
    } as any);
    expect(wrapper).toBeTruthy()
  })
 it("renders as expected", () => {
  const wrapper = shallowMount(EditProfilePage, {
    global: {
      plugins: [store],
    },
  } as any);
    expect(wrapper.html()).toMatchSnapshot()
  })
})
