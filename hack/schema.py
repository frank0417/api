#!/usr/bin/env python

import fnmatch
import json
import os
import re
import subprocess
import sys
from os.path import expandvars

# binary -> GOOS -> GOARCH
# windows	386
# windows	amd64

# Debian package
# https://gist.github.com/rcrowley/3728417

API_ROOT = expandvars("$GOPATH/src/github.com/appscode/api")
VALID_FORMATS = ['date-time',
                 'email',
                 'hostname',
                 'ipv4',
                 'ipv6',
                 'uri']
PROTO_PKG_REGEX = ur'^package\s*(?P<pkg>[^;]*);$'
GO_PKG_REGEX = ur'^option go_package\s*=\s*"(?P<pkg>[^;]*)"\s*;$'


def call(cmd, stdin=None, cwd=API_ROOT):
    print(cmd)
    subprocess.call([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)


def write_file(name, content):
    dir = os.path.dirname(name)
    if not os.path.exists(dir):
        os.makedirs(dir)
    with open(name, 'w') as f:
        return f.write(content)


def append_file(name, content):
    with open(name, 'a') as f:
        return f.write(content)


# TODO: use unicode encoding
def read_json(name):
    try:
        with open(name, 'r') as f:
            return json.load(f)
    except IOError:
        return {}


def write_json(obj, name):
    with open(name, 'w') as f:
        return json.dump(obj, f, sort_keys=True, indent=2, separators=(',', ': '))


def deps():
    # specify git commit_hash to pin to specific version
    get_pkgs = [
        {
            'pkg': 'github.com/golang/protobuf/protoc-gen-go',
            'commit_hash': '888eb0692c857ec880338addf316bd662d5e630e',
            'install': True
        },
        {
            'pkg': 'google.golang.org/grpc',
            'commit_hash': '0032a855ba5c8a3c8e0d71c2deef354b70af1584'
        },
        {
            'pkg': 'github.com/golang/glog',
            'commit_hash': '44145f04b68cf362d9c4df2182967c2275eaefed'
        },
        {'pkg': 'github.com/jteeuwen/go-bindata/...'},
        {'pkg': 'golang.org/x/tools/cmd/goimports'},
        {'pkg': 'github.com/xeipuuv/gojsonschema'},
    ]
    for cfg in get_pkgs:
        call('go get -u ' + cfg['pkg'])
        if cfg.get('commit_hash', ''):
            call('git checkout ' + cfg['commit_hash'], cwd=expandvars('$GOPATH/src/' + cfg['pkg'].rstrip('/...')))
        if cfg.get('install', False):
            call('go install ./...', cwd=expandvars('$GOPATH/src/' + cfg['pkg'].rstrip('/...')))
    # special treatment for grc-gateway
    call('rm -rf $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway')
    call('mkdir -p $GOPATH/src/github.com/grpc-ecosystem')
    call('git clone https://github.com/appscode/grpc-gateway.git', cwd=expandvars('$GOPATH/src/github.com/grpc-ecosystem'))
    call('go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway')
    call('go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger')
    # Copy google http apis proto files
    call('mkdir -p $GOPATH/src/github.com/google')
    call('git clone https://github.com/googleapis/googleapis.git', cwd=expandvars('$GOPATH/src/github.com/google'))


def gen_assets():
    call('go get github.com/jteeuwen/go-bindata/...')
    call('go-bindata -ignore=\\.go -o meta/data.go -pkg meta meta/...')


def fix_swagger_schema():
    for root, dirnames, filenames in os.walk(API_ROOT):
        for filename in fnmatch.filter(filenames, '*.swagger.json'):

            rel_path = root[len(API_ROOT) + 1:]
            parts = rel_path.split('/', 2)
            if len(parts) != 2:
                continue
            # api_name = parts[0]
            # api_version = parts[1]

            swagger = os.path.join(root, filename)
            spec = read_json(swagger)
            spec["basePath"] = "/"
            spec["host"] = "api.appscode.com"
            spec["info"]["version"] = parts[1][1:]
            spec["schemes"] = ["https"]
            write_json(spec, swagger)


def swagger_defs(defs):
    stack = []
    result = {
        'requests': {},
        'responses': {}
    }
    for key in defs.keys():
        if key.endswith("Request"):
            stack.append(key)
        elif key.endswith("Response"):
            result['responses'][key] = defs[key]
    while stack:
        name = stack.pop()
        schema = defs[name]
        result['requests'][name] = schema
        if 'properties' in schema:
            for p, v in schema['properties'].iteritems():
                if '$ref' in v:
                    nw_obj = v['$ref'][len('#/definitions/'):]
                    if nw_obj not in result['requests']:
                        stack.append(nw_obj)
                if 'format' in v and not v['format'] in VALID_FORMATS:
                    v.pop('format', None)
                if 'additionalProperties' in v and 'format' in v['additionalProperties'] and not v['additionalProperties']['format'] in VALID_FORMATS:
                    v['additionalProperties'].pop('format', None)
                if 'items' in v:
                    if '$ref' in v['items']:
                        nw_obj = v['items']['$ref'][len('#/definitions/'):]
                        if nw_obj not in result['requests']:
                            stack.append(nw_obj)
                    if 'format' in v['items'] and not v['items']['format'] in VALID_FORMATS:
                        v['items'].pop('format', None)
    return result


def generate_json_schema():
    call("find . | grep schema.json | xargs rename 's/schema.json/schema.json.ext/' {}")
    for root, dirnames, filenames in os.walk(API_ROOT):
        for filename in fnmatch.filter(filenames, '*.swagger.json'):
            swagger = os.path.join(root, filename)
            schema = os.path.join(root, filename.replace('.swagger.', '.schema.'))
            ext_schema = os.path.join(root, filename.replace('.swagger.json', '.schema.json.ext'))
            gen_defs = swagger_defs(read_json(swagger)['definitions'])['requests']
            if os.path.exists(ext_schema):
                # merge
                ext_defs = read_json(ext_schema)['definitions']
                for m, mspec in gen_defs.iteritems():
                    if 'properties' in mspec.keys():
                        for f, fspec in mspec['properties'].iteritems():
                            if f in [
                                'auth_secret_name',
                                'bucket_name',
                                'cloud_credential',
                                'cluster_name',
                                'name',
                                'namespace',
                                'secret_name',
                                'service_name',
                                'snapshot_name'
                            ]:
                                if 'maxLength' not in fspec:
                                    fspec['maxLength'] = 63
                                if 'pattern' not in fspec:
                                    fspec['pattern'] = "^[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?$"

                            if m in ext_defs \
                                    and 'properties' in ext_defs[m] \
                                    and f in ext_defs[m]['properties'] \
                                    and set(fspec.keys()) != set(ext_defs[m]['properties'][f].keys()):
                                print mspec['properties'][f]
                                print ext_defs[m]['properties'][f]
                                mspec['properties'][f] = ext_defs[m]['properties'][f]
            write_json({'definitions': gen_defs}, schema)
    call("(find . | grep schema.json.ext | xargs rm) || true")


def schema_go(pkg, defs):
    result = {
        'requests': {},
        'responses': {}
    }
    for key in defs['requests'].keys():
        if key.startswith(pkg) and key.endswith("Request"):
            schema = defs['requests'][key]
            result['requests'][key[len(pkg):]] = schema
            dep_defs = {}
            stack = []
            stack.append(schema)
            while stack:
                sch = stack.pop()
                if 'properties' in sch:
                    for p, v in sch['properties'].iteritems():
                        if '$ref' in v:
                            nw_obj = v['$ref'][len('#/definitions/'):]
                            if nw_obj not in dep_defs:
                                dep_defs[nw_obj] = defs['requests'][nw_obj]
                                stack.append(defs['requests'][nw_obj])
                        if 'items' in v and '$ref' in v['items']:
                            nw_obj = v['items']['$ref'][len('#/definitions/'):]
                            if nw_obj not in dep_defs:
                                dep_defs[nw_obj] = defs['requests'][nw_obj]
                                stack.append(defs['requests'][nw_obj])
            if dep_defs:
                schema['definitions'] = dep_defs
            schema['$schema'] = 'http://json-schema.org/draft-04/schema#'
    for key in defs['responses'].keys():
        print key
        if key.startswith(pkg) and key.endswith("Response"):
            schema = defs['responses'][key]
            result['responses'][key[len(pkg):]] = schema
    return result


def render_schema_go(pkg, schemas):
    contents = """package {}

// Auto-generated. DO NOT EDIT.
""".format(pkg)

    imports = []
    if schemas['requests']:
        imports.append("github.com/xeipuuv/gojsonschema")
        imports.append("github.com/golang/glog")
    if schemas['responses']:
        imports.append("github.com/appscode/api/dtypes")
    imports.sort()
    if imports:
        contents += 'import (\n'
        for pkg in imports:
            contents += '	"{}"\n'.format(pkg)
        contents += ')\n\n'

    for key in schemas['requests'].keys():
        contents += 'var {0}Schema *gojsonschema.Schema\n'.format(key[0:1].lower() + key[1:])
    contents += '\n'

    if schemas['requests']:
        contents += """func init() {
	var err error
"""
        for key, sch in schemas['requests'].iteritems():
            var_name = key[0:1].lower() + key[1:]
            sch_str = json.dumps(sch, sort_keys=True, indent=2, separators=(',', ': '))
            sch_str = sch_str.replace('`', '` + "`" + `')
            contents += '	{0}Schema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{1}`))\n'.format(
                var_name, sch_str)
            contents += """	if err != nil {
		glog.Fatal(err)
	}
"""
        contents += '}\n\n'

    for key in schemas['requests'].keys():
        contents += 'func (m *' + key + ') IsValid() (*gojsonschema.Result, error) {\n'
        contents += '	return {}Schema.Validate(gojsonschema.NewGoLoader(m))\n'.format(key[0:1].lower() + key[1:])
        contents += '}\n'
        contents += 'func (m *' + key + ') IsRequest() {}\n\n'

    for key in schemas['responses'].keys():
        contents += 'func (m *' + key + ') SetStatus(s *dtypes.Status) {\n'
        contents += '	m.Status = s\n'
        contents += '}\n'

    return contents


def detect_schema_pkg(swagger):
    proto = swagger[:-len('.swagger.json')] + '.proto'
    with open(proto) as f:
        content = f.read()
        m = re.search(PROTO_PKG_REGEX, content, re.MULTILINE)
        if m:
            pkg = m.group('pkg')
            parts = pkg.split(".")
            for x in range(0, len(parts)):
                prefix = str.join("", parts[x:])
                for key, defs in read_json(swagger)['definitions'].iteritems():
                    if key.startswith(prefix) and key.endswith("Request"):
                        return prefix
        else:
            print("Failed to detect package.")
            sys.exit(1)


def detect_go_pkg(swagger):
    proto = swagger[:-len('.swagger.json')] + '.proto'
    with open(proto) as f:
        content = f.read()
        m = re.search(GO_PKG_REGEX, content, re.MULTILINE)
        if m:
            return m.group('pkg')
        else:
            m = re.search(PROTO_PKG_REGEX, content, re.MULTILINE)
            if m:
                return m.group('pkg')
            else:
                print("Failed to detect package.")
                sys.exit(1)


def generate_go_schema():
    for root, dirnames, filenames in os.walk(API_ROOT):
        for filename in fnmatch.filter(filenames, '*.swagger.json'):
            swagger = os.path.join(root, filename)
            schema = os.path.join(root, filename.replace('.swagger.', '.schema.'))
            go = schema[:-len('.json')] + '.go'
            print go
            pkg = detect_schema_pkg(swagger)
            if pkg:
                defs = swagger_defs(read_json(swagger)['definitions'])
                # overwrite requests with json schema from *.schema.json
                # to preserve hand written rules
                defs['requests'] = read_json(schema)['definitions']
                schemas = schema_go(pkg, defs)
                if schemas['requests'] or schemas['responses']:
                    write_file(go, render_schema_go(detect_go_pkg(swagger), schemas))


def apply_naming_policy():
    for root, dirnames, filenames in os.walk(API_ROOT):
        for filename in fnmatch.filter(filenames, '*.schema.json'):
            schema = os.path.join(root, filename)
            print schema
            content = read_json(schema)
            for key, defs in content['definitions'].iteritems():
                if 'properties' in defs:
                    for p, v in defs['properties'].iteritems():
                        if p in [
                            'cluster_name',
                            'namespace', 'name',
                            'bucket_name',
                            'secret_name',
                            'snapshot_name',
                            'auth_secret_name',
                            'cloud_credential'
                        ]:
                            print '====>>>> ' + p
                            if 'maxLength' not in v:
                                v['maxLength'] = 63
                            if 'pattern' not in v:
                                v['pattern'] = "^[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?$"
            write_json(content, schema)


if __name__ == "__main__":
    if len(sys.argv) > 1:
        # http://stackoverflow.com/a/834451
        # http://stackoverflow.com/a/817296
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        generate_json_schema()
        # apply_naming_policy()
        generate_go_schema()
