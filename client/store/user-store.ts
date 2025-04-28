import { atomWithStorage } from "jotai/utils";

import { User } from "@/types/user";

// Create atom with storage to persist user data across page refreshes
export const userAtom = atomWithStorage<User | null>("user", null);
